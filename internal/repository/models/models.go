package models

type Item struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	PubDate     string `json:"pubDate"`
}
