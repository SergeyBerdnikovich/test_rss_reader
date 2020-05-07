package rssreader

import (
	"context"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

type reader struct{}

type parseResult struct {
	URL  string
	Feed *gofeed.Feed
	Err  error
}

// Parse is a function which parses asynchronously rss urls
// As arguments there are context and slice of urls
// Retruns slice of RssItem and the concatinated errors if the are exists
func Parse(ctx context.Context, urls []string) ([]RssItem, error) {
	return reader{}.parse(ctx, urls)
}

func (r reader) parse(ctx context.Context, urls []string) ([]RssItem, error) {
	parseResults := make(chan parseResult)

	for _, url := range urls {
		url := url

		go parseFeed(ctx, url, parseResults)
	}

	rssItems := []RssItem{}
	parseErrs := []string{}

	for i := 0; i < len(urls); i++ {
		parseResult := <-parseResults

		if parseResult.Err != nil {
			parseErrs = append(parseErrs, errors.Wrap(parseResult.Err, parseResult.URL).Error())
		} else {
			items := parseFeedItems(*parseResult.Feed)
			rssItems = append(rssItems, items...)
		}
	}

	if len(parseErrs) != 0 {
		return rssItems, joinErrors(parseErrs)
	}

	return rssItems, nil
}

func parseFeed(ctx context.Context, url string, parseResults chan<- parseResult) {
	feed, err := gofeed.NewParser().ParseURLWithContext(url, ctx)
	parseResults <- parseResult{URL: url, Feed: feed, Err: err}
}

func parseFeedItems(feed gofeed.Feed) []RssItem {
	items := make([]RssItem, 0, len(feed.Items))

	for _, item := range feed.Items {
		items = append(items, RssItem{
			Title:       item.Title,
			Source:      feed.Title,
			SourceURL:   feed.Link,
			Link:        item.Link,
			PublishDate: item.PublishedParsed,
			Description: item.Description,
		})
	}

	return items
}

func joinErrors(errsMessages []string) error {
	return errors.New(strings.Join(errsMessages, ", "))
}
