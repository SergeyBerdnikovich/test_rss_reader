package rssreader

import "time"

// RssItem is an entity which describes rss feed item
type RssItem struct {
	Title       string
	Source      string
	SourceURL   string
	Link        string
	PublishDate time.Time
	Description string
}
