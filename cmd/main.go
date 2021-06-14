package main

import (
	"fmt"
	loaderpkg "github.com/keithzetterstrom/araneus/internal/loader"
	"github.com/keithzetterstrom/araneus/internal/models"
	parserpkg "github.com/keithzetterstrom/araneus/internal/parser"
	"github.com/keithzetterstrom/araneus/internal/repository/elasticsearch"
)

const parseURL = "https://lenta.ru/rss/news"
const (
	search = "люди"
	termAggregation = "author"
	cardinalityAggregation = "author"
)

func main()  {
	fmt.Println("araneus")

	loader := loaderpkg.NewLoader()

	parser := parserpkg.NewParser(loader)

	es, err := elasticsearch.NewESClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = es.CreateIndex()
	if err != nil {
		fmt.Println(err)
		return
	}

	existsItems, err := es.GetLastItems()
	if err != nil {
		fmt.Println(err)
		return
	}

	newItems, err := parser.ParseItemsPage(parseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	items := findUniqItems(existsItems, newItems)

	itemsWithText, err := parser.ParseItemPage(items)
	if err != nil {
		fmt.Println(err)
		return
	}


	err = es.InsertItems(itemsWithText)
	if err != nil {
		fmt.Println(err)
		return
	}

	itemsByKey, err := es.FullTextSearch(search)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("\nFind by word, ", search)
	PrintItems(itemsByKey)

	termAggr, err := es.TermAggregationByField(termAggregation)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("\nTerm aggregation by %s\n", termAggregation)
	for _, t := range termAggr {
		fmt.Printf("%s: %v\n", termAggregation, t.Key)
		fmt.Printf("Count: %d\n", t.Count)
	}

	cardinalityAggr, err := es.CardinalityAggregationByField(cardinalityAggregation)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("\nCardinality aggregation by %s: %.0f\n", cardinalityAggregation, cardinalityAggr)

	dataHistogram, err := es.DateHistogramAggregation()
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("\nDate histogram aggregation")
	for _, t := range dataHistogram {
		fmt.Printf("Key: %v\n", t.Key)
		fmt.Printf("Count: %d\n", t.Count)
	}
}

func PrintItems(items []models.Item)  {
	for _, item := range items{
		fmt.Println("Title: " + item.Title)
		fmt.Println("PubDate: " + item.PubDate.String())
		fmt.Println("Author: " + item.Author)
		fmt.Println("Link: " + item.Link)
	}
}

func findUniqItems(existsItems []*models.Item, newItems []*models.Item) []*models.Item {
	existsItemsMap := map[string]*models.Item{}
	var uniqItems []*models.Item

	for _, item := range existsItems {
		existsItemsMap[item.Link] = item
	}

	for _, newItem := range newItems {
		if _, ok := existsItemsMap[newItem.Link]; !ok {
			uniqItems = append(uniqItems, newItem)
		}
	}

	return uniqItems
}
