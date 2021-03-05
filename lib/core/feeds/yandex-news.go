package feeds

import (
	"encoding/xml"
	"time"
)

// YandexNewsXML todo
type YandexNewsXML struct {
	XMLName         xml.Name `xml:"rss"`
	Channel         *ynFeed  `xml:"channel"`
	YandexNamespace string   `xml:"xmlns:yandex,attr"`
	MediaNamespace  string   `xml:"xmlns:media,attr"`
	Version         string   `xml:"version,attr"`
}

type ynFeed struct {
	Base          string     `xml:"xml:base,attr"`
	Description   string     `xml:"description"`
	Images        []*ynImage `xml:"yandex:logo"`
	Items         []*ynItem  `xml:"item"`
	Language      string     `xml:"language,omitempty"`
	LastBuildDate string     `xml:"lastBuildDate,omitempty"`
	Link          string     `xml:"link"`
	Title         string     `xml:"title"`
}

type ynImage struct {
	Typeof string `xml:"typeof,attr,omitempty"`
	Value  string `xml:",chardata"`
}

type ynItem struct {
	Author      string        `xml:"author"`
	Title       string        `xml:"title"`
	Link        string        `xml:"link"`
	Content     string        `xml:"yandex:full-text"`
	Description string        `xml:"description"`
	PublishedAt string        `xml:"pubDate"`
	PDALink     string        `xml:"pdalink"`
	Category    string        `xml:"category"`
	GUID        *ynItemGUID   `xml:"guid"`
	Enclosure   *rssEnclosure `xml:"enclosure"`
}

type ynItemGUID struct {
	IsPermaLink bool   `xml:"isPermalink,attr"`
	Value       string `xml:",chardata"`
}

type ynItemEnclosure struct {
	URL string `xml:"url,attr"`
}

// YandexNews todo
type YandexNews struct {
	*GenericFeed
}

// NewYandexNews todo
func NewYandexNews(genericFeed *GenericFeed) *YandexNews {
	return &YandexNews{genericFeed}
}

// ToXML todo
func (yn *YandexNews) ToXML() interface{} {
	srcXML := &ynFeed{
		Base:        yn.Base,
		Title:       yn.Title,
		Link:        yn.Link.Href,
		Description: yn.Description,
	}

	if yn.Image != nil {
		srcXML.Images = append(srcXML.Images, &ynImage{
			Value: yn.Image.Link,
		})
		srcXML.Images = append(srcXML.Images, &ynImage{
			Typeof: "square",
			Value:  yn.Image.Link,
		})
	}

	for _, i := range yn.Items {
		item := &ynItem{
			Title:       i.Title,
			Category:    i.Category,
			Description: i.Description,
			GUID:        &ynItemGUID{IsPermaLink: true, Value: i.Link.Href},
			Link:        i.Link.Href,
			PDALink:     i.Link.Href,
			PublishedAt: i.Created.Format(time.RFC1123Z),
		}

		if len(i.Content) > 0 {
			item.Content = i.GetPlainBody()
		}

		if i.Attachment != nil && i.Attachment.Type != "" {
			item.Enclosure = &rssEnclosure{
				URL:    i.Attachment.URL,
				Type:   i.Attachment.Type,
				Length: i.Attachment.Length,
			}
		}

		if len(i.Authors) > 0 {
			item.Author = i.Authors[0].Name
		}

		srcXML.Items = append(srcXML.Items, item)
	}

	return &YandexNewsXML{
		Channel:         srcXML,
		MediaNamespace:  "http://search.yahoo.com/mrss/",
		YandexNamespace: "http://news.yandex.ru",
		Version:         RssVersion2,
	}
}
