package parser

import (
	"encoding/xml"
	"github.com/keithzetterstrom/araneus/internal/models"
	logtool "github.com/keithzetterstrom/araneus/tools/logger"
)

type parser struct {
	Logger logtool.Logger
	ByteValue []byte
}

type Parser interface {
	ParsePage() ([]models.Item, error)
}

func NewParser(byteValue []byte, logger logtool.Logger) Parser {
	return &parser{
		Logger: logger,
		ByteValue: byteValue,
	}
}

func (p * parser) ParsePage() ([]models.Item, error) {
	var channel RSS

	err := xml.Unmarshal(p.ByteValue, &channel)
	if err != nil {
		p.Logger.ErrorLogger.Println(err)
		return nil, err
	}

	var items []models.Item

	for _, item := range channel.Channel.Items {
		items = append(items, convertXMLItemToItem(item))
	}
	return items, nil
}

func convertXMLItemToItem(item Item) models.Item {
	return models.Item{
		Guid:        item.Guid,
		Author:      item.Author,
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		PubDate:     item.PubDate,
		Category:    item.Category,
	}
}
