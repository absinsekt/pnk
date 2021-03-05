package feeds

import (
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

// GenericFeed todo
type GenericFeed struct {
	Author      *Author
	Base        string
	Copyright   string
	Created     time.Time
	Description string
	ID          string
	Image       *Image
	Items       []*Item
	Language    string
	Link        *Link
	Subtitle    string
	Title       string
	Updated     time.Time
}

// Add todo
func (f *GenericFeed) Add(item *Item) {
	f.Items = append(f.Items, item)
}

// Attachment todo
type Attachment struct {
	Type   string
	Length uint64
	URL    string
}

// Author todo
type Author struct {
	Email string
	Name  string
}

// Image todo
type Image struct {
	Height uint64
	Link   string
	Title  string
	URL    string
	Width  uint64
}

// Item todo
type Item struct {
	Authors     []*Author
	Category    string
	Content     string
	Created     time.Time
	Description string
	Attachment  *Attachment
	ID          string
	Link        *Link
	Source      *Link
	Title       string
	Updated     time.Time
}

// GetPlainBody todo
func (i *Item) GetPlainBody() string {
	sanitizer := bluemonday.StripTagsPolicy()
	result := strings.ReplaceAll(i.Content, "\r\n", " ")

	return sanitizer.Sanitize(result)
}

// Link todo
type Link struct {
	Href string
	Rel  string
}
