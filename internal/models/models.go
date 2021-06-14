package models

import "time"

type Item struct {
	Guid        string    `json:"guid"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"pub_date"`
	Category    string    `json:"category"`
	Text        string    `json:"text"`
}
