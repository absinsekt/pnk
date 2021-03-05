package feeds

import (
	"encoding/xml"
	"time"
)

const (
	RssVersion1 = "1.0"
	RssVersion2 = "2.0"
)

// RssXml todo
type RssXml struct {
	XMLName          xml.Name `xml:"rss"`
	Channel          *rssFeed `xml:"channel"`
	ContentNamespace string   `xml:"xmlns:content,attr"`
	Version          string   `xml:"version,attr"`
}

type rssFeed struct {
	Description   string     `xml:"description"`
	Image         *rssImage  `xml:"image"`
	Items         []*rssItem `xml:"item"`
	Language      string     `xml:"language,omitempty"`
	LastBuildDate string     `xml:"lastBuildDate,omitempty"`
	Link          string     `xml:"link"`
	Title         string     `xml:"title"`
}

type rssImage struct {
	Height uint64 `xml:"height,omitempty"`
	Link   string `xml:"link"`
	Title  string `xml:"title,omitempty"`
	URL    string `xml:"url,omitempty"`
	Width  uint64 `xml:"width,omitempty"`
}

type rssItem struct {
	Authors     []string      `xml:"author,omitempty"`
	Category    string        `xml:"category,omitempty"`
	Content     *rssContent   `xml:"content:encoded"`
	Description string        `xml:"description"`
	Enclosure   *rssEnclosure `xml:"enclosure"`
	GUID        string        `xml:"guid,omitempty"`
	Link        string        `xml:"link"`
	PubDate     string        `xml:"pubDate,omitempty"`
	Title       string        `xml:"title"`
}

type rssContent struct {
	Content string `xml:",cdata"`
}

type rssEnclosure struct {
	Length uint64 `xml:"length,attr"`
	Type   string `xml:"type,attr"`
	URL    string `xml:"url,attr"`
}

// Rss todo
type Rss struct {
	*GenericFeed
	version string
}

// NewRss todo
func NewRss(genericFeed *GenericFeed, version string) *Rss {
	return &Rss{genericFeed, version}
}

// ToXML todo
func (rss *Rss) ToXML() interface{} {
	var result *RssXml

	srcXML := &rssFeed{
		Title:         rss.Title,
		Description:   rss.Description,
		Items:         []*rssItem{},
		Language:      rss.Language,
		LastBuildDate: rss.Updated.Format(time.RFC1123Z),
		Link:          rss.Link.Href,
	}

	if rss.Image != nil {
		srcXML.Image = &rssImage{
			Height: rss.Image.Height,
			Link:   rss.Image.Link,
			Title:  rss.Image.Title,
			URL:    rss.Image.URL,
			Width:  rss.Image.Width,
		}
	}

	for _, i := range rss.Items {
		item := &rssItem{
			Category:    i.Category,
			Description: i.Description,
			GUID:        i.ID,
			Link:        i.Link.Href,
			PubDate:     i.Updated.Format(time.RFC1123Z),
			Title:       i.Title,
		}

		if len(i.Content) > 0 {
			item.Content = &rssContent{Content: i.Content}
		}

		if i.Attachment != nil && i.Attachment.Type != "" && i.Attachment.Length != 0 {
			item.Enclosure = &rssEnclosure{
				URL:    i.Attachment.URL,
				Type:   i.Attachment.Type,
				Length: i.Attachment.Length,
			}
		}

		for _, a := range i.Authors {
			item.Authors = append(item.Authors, a.Name)
		}

		srcXML.Items = append(srcXML.Items, item)
	}

	switch rss.version {
	default:
		result = &RssXml{
			Channel:          srcXML,
			ContentNamespace: "http://purl.org/rss/1.0/modules/content/",
			Version:          rss.version,
		}
	}

	return result
}
