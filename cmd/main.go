package main

import (
	"fmt"
	loaderpkg "github.com/keithzetterstrom/araneus/internal/loader"
	"github.com/keithzetterstrom/araneus/internal/models"
	parserpkg "github.com/keithzetterstrom/araneus/internal/parser"
	logtool "github.com/keithzetterstrom/araneus/tools/logger"
	"log"
)

var parsURL = "https://lenta.ru/rss/news"

func main()  {
	fmt.Println("araneus")

	logFile, err := logtool.OpenLogFile()
	if err != nil {
		log.Fatal()
	}

	logger := logtool.NewLogger(logFile)

	loader := loaderpkg.NewLoader(logger)

	res, err := loader.LoadPage(parsURL)
	if err != nil {
		return
	}

	parser := parserpkg.NewParser(res, logger)

	items, err := parser.ParsePage()
	if err != nil {
		return
	}

	PrintItems(items)
}

func PrintItems(items []models.Item)  {
	for _, item := range items{
		fmt.Println("Description: " + item.Description)
		fmt.Println("Title: " 		+ item.Title)
		fmt.Println("PubDate: " 	+ item.PubDate)
		fmt.Println("Author: " 		+ item.Author)
		fmt.Println("Link: " 		+ item.Link)
		fmt.Println("Category: " 	+ item.Category)
	}
}
