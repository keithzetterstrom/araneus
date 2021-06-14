package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	loaderpkg "github.com/keithzetterstrom/araneus/internal/loader"
	"github.com/keithzetterstrom/araneus/internal/models"
	"time"
)

const nameSelector = ".js-topic__text"

type parser struct {
	Loader loaderpkg.Loader
}

type Parser interface {
	ParseItemsPage(url string) ([]*models.Item, error)
	ParseItemPage([]*models.Item) ([]*models.Item, error)
}

func NewParser(loader loaderpkg.Loader) Parser {
	return &parser{
		Loader: loader,
	}
}

func (p * parser) ParseItemsPage(url string) ([]*models.Item, error) {
	var channel RSS

	res, err := p.Loader.LoadPage(url)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(res, &channel)
	if err != nil {
		return nil, err
	}

	var items []*models.Item

	for _, item := range channel.Channel.Items {
		i, err := convertXMLItemToItem(item)
		if err != nil {
			continue
		}
		items = append(items, i)
	}
	return items, nil
}

func convertXMLItemToItem(item Item) (*models.Item, error) {
	date, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		return nil, err
	}
	return &models.Item{
		Guid:        item.Guid,
		Author:      item.Author,
		Title:       item.Title,
		Link:        item.Link,
		Description: item.Description,
		PubDate:     date,
		Category:    item.Category,
	}, nil
}

func (p * parser) ParseItemPage(items []*models.Item) ([]*models.Item, error) {
	for _, item := range items {
		text, err := p.parseItem(item.Link)
		if err != nil {
			fmt.Println(err)
			continue
		}
		item.Text = text
		break
	}
	return items, nil
}

func (p * parser) parseItem(url string) (string, error) {
	res, err := p.Loader.LoadPage(url)
	if err != nil {
		return "", err
	}

	page, err := goquery.NewDocumentFromReader(bytes.NewReader(res))
	if err != nil {
		return "", err
	}

	text := ""
	page.Find(nameSelector).Each(func(i int, s *goquery.Selection) {
		text += s.Find("p").Text()
	})

	return text, nil
}
