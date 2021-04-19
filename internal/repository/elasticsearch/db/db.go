package db

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	internalmodels "github.com/keithzetterstrom/araneus/internal/models"
	"github.com/keithzetterstrom/araneus/internal/repository/models"
	"github.com/keithzetterstrom/araneus/tools/hash"
	"log"
	"sync"
)

type ElasticSearch struct {
	clientES *elasticsearch.Client
}

func NewESClient() (*ElasticSearch, error) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	return &ElasticSearch{clientES: es}, nil
}

func (es *ElasticSearch) GetInfo() {
	res, err := es.clientES.Info()
	if err != nil {
		fmt.Printf("Error getting response: %s \n", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Printf("Error: %s \n", res.String())
	}

	r := map[string]interface{}{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Printf("Error parsing the response body: %s \n", err)
	}

	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
}

func (es *ElasticSearch) CreateESIndex(internalItems []internalmodels.Item) error {
	var wg sync.WaitGroup

	for _, internalItem := range internalItems {
		wg.Add(1)

		item := convertItemToRepositoryItem(internalItem)

		go es.writeItem(item, &wg)
	}

	wg.Wait()

	return nil
}

func (es *ElasticSearch) writeItem(item models.Item, wg * sync.WaitGroup)  {
	defer wg.Done()

	b, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	id := hash.MD5(item)

	req := esapi.IndexRequest{
		Index:      "items",
		DocumentID: id,
		Body:       bytes.NewReader(b),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es.clientES)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		fmt.Println(err)
		return
	}
}

func convertItemToRepositoryItem(item internalmodels.Item) models.Item {
	return models.Item{
		Title: item.Title,
		Author: item.Author,
		Description: item.Description,
		PubDate: item.PubDate,
		Link: item.Link,
	}
}
