package elasticsearch

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/keithzetterstrom/araneus/internal/models"
	"github.com/olivere/elastic/v7"
	"os"

	"log"
)

const indexName = "items"

const (
	dateFormat   = "dd.MM.YYYY"
	dateInterval = "1d"
)

type ESClient struct {
	client *elastic.Client
}

func NewESClient() (ESClient, error) {
	client, err := elastic.NewSimpleClient(elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)))
	if err != nil {
		return ESClient{}, err
	}

	return ESClient{client: client}, nil
}

func (es *ESClient) CreateIndex() error {
	ctx := context.Background()

	exists, err := es.client.
		IndexExists(indexName).
		Do(ctx)
	if err != nil {
		log.Println("Error with index exists request", err)
	}

	if exists {
		return nil
	}

	_, err = es.client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		log.Println("Error with create index", err)
		return err
	}

	err = es.createMapping()
	if err != nil {
		return err
	}

	return nil
}

func (es *ESClient) createMapping() error {
	ctx := context.Background()

	_, err := es.client.IndexPutSettings().Index(indexName).BodyString(PutSetting).Do(ctx)
	if err != nil {
		log.Println("Error with put mapping", err)
		return err
	}
	_, err = es.client.PutMapping().
		Index(indexName).
		BodyString(PutMapping).
		Do(ctx)
	if err != nil {
		log.Println("Error with put mapping", err)
		return err
	}

	return nil
}

func (es *ESClient) InsertItems(items []*models.Item) error {
	ctx := context.Background()

	for _, item := range items {
		_, err := es.client.Index().
			Index(indexName).
			Id(item.Link).
			BodyJson(&item).
			Do(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (es *ESClient) GetLastItems() ([]*models.Item, error) {
	ctx := context.Background()

	query := elastic.NewMatchAllQuery()

	res, err := es.client.Search().
		Index(indexName).
		Query(query).
		Size(100).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	var items []*models.Item

	for _, hit := range res.Hits.Hits {
		var t *models.Item
		err := json.Unmarshal(hit.Source, &t)
		if err != nil {
			return nil, err
		}

		items = append(items, t)
	}

	return items, nil
}

func (es *ESClient) FullTextSearch(key string) ([]models.Item, error) {
	ctx := context.Background()

	queryString := elastic.NewQueryStringQuery(key)

	res, err := es.client.Search().
		Index(indexName).
		Query(queryString).
		Do(ctx)
	if err != nil {
		return []models.Item{}, err
	}

	if res.Hits.TotalHits.Value == 0 {
		return []models.Item{}, errors.New("Not found ")
	}

	var items []models.Item

	for _, hit := range res.Hits.Hits {
		var t models.Item
		err := json.Unmarshal(hit.Source, &t)
		if err != nil {
			return nil, err
		}

		items = append(items, t)
	}

	return items, nil
}

func (es *ESClient) TermAggregationByField(field string) ([]AggregationOutput, error) {
	ctx := context.Background()

	aggregationQuery := elastic.NewTermsAggregation().
		Field(addKeyWord(field)).
		Size(30).
		OrderByCountDesc()

	result, err := es.client.Search().
		Index(indexName).
		Aggregation(indexName, aggregationQuery).
		Do(ctx)
	if err != nil {
		e, ok := err.(*elastic.Error)
		if ok {
			log.Printf("Got error from elastic: %s", e.Details)
		}
		return []AggregationOutput{}, err
	}

	rawMsg := result.Aggregations[indexName]

	ar := elastic.AggregationBucketKeyItems{}

	err = json.Unmarshal(rawMsg, &ar)
	if err != nil {
		return nil, err
	}

	var termsAggregations []AggregationOutput

	for _, item := range ar.Buckets {
		termsAggregations = append(termsAggregations, AggregationOutput{
			Key:   item.Key,
			Count: item.DocCount,
		})
	}
	return termsAggregations, nil
}

func (es *ESClient) CardinalityAggregationByField(field string) (float64, error) {
	ctx := context.Background()

	aggregationQuery := elastic.NewCardinalityAggregation().Field(addKeyWord(field))

	result, err := es.client.Search().
		Index(indexName).
		Aggregation(indexName, aggregationQuery).
		Do(ctx)
	if err != nil {
		e, ok := err.(*elastic.Error)
		if ok {
			log.Printf("Got error from elastic %s", e.Details)
		}
		return 0, err
	}

	rawMsg := result.Aggregations[indexName]

	var ar elastic.AggregationValueMetric

	err = json.Unmarshal(rawMsg, &ar)
	if err != nil {
		return 0, err
	}

	return *ar.Value, nil
}

func (es *ESClient) DateHistogramAggregation() ([]AggregationOutput, error) {
	ctx := context.Background()
	dailyAggregation := elastic.NewDateHistogramAggregation().
		Field("pub_date").
		CalendarInterval(dateInterval).
		Format(dateFormat)

	result, err := es.client.Search().
		Index(indexName).
		Aggregation(indexName, dailyAggregation).
		Do(ctx)
	if err != nil {
		e, ok := err.(*elastic.Error)
		if ok {
			log.Printf("Got error from elastic %s", e.Details)
		}
		return []AggregationOutput{}, err
	}

	hist, found := result.Aggregations.Histogram(indexName)
	if !found {
		return []AggregationOutput{}, errors.New("Not found ")
	}

	var dateHistogramAggregations []AggregationOutput

	for _, bucket := range hist.Buckets {
		dateHistogramAggregations = append(dateHistogramAggregations, AggregationOutput{
			Key:   *bucket.KeyAsString,
			Count: bucket.DocCount,
		})
	}

	return dateHistogramAggregations, nil
}

func addKeyWord(field string) string {
	if field == "title" || field == "description" {
		return field + ".keyword"
	}

	return field
}
