package rssreader

import "time"

// RssItem is an entity which describes rss feed item
type RssItem struct {
	Title       string     `json:"title"`
	Source      string     `json:"source"`
	SourceURL   string     `json:"source_url"`
	Link        string     `json:"link"`
	PublishDate *time.Time `json:"publish_date"`
	Description string     `json:"description"`
}
