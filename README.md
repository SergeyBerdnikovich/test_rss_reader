# test_rss_reader
a test library for parsing rss for https://www.emerchantpay.com/

<h2>Basic Usage:</h2>
<p>
import <code>rssreader</code> package and call <code>Parse</code> function:
</br>
<pre>
import "github.com/SergeyBerdnikovich/test_rss_reader/pkg/rssreader"

posts, err := rssreader.Parse(context.Background(), []string{"http://feeds.twit.tv/twit.xml"})
</pre>
</p>
<p>
where <code>posts</code> is a slice of rss items:
</br>
<pre>
type RssItem struct {
Title       string
Source      string
SourceURL   string
Link        string
PublishDate time.Time
Description string
}
</pre>
</p>
and <code>err</code> is a concatenated <code>error</code> or <code>nil</code>
