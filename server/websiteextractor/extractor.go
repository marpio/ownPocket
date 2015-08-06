package contentextractor

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/kennygrant/sanitize"
	"github.com/mauidude/go-readability"
)

// Content of a wesite
type Content struct {
	url              string
	title            string
	readableContent  string
	sanitizedContent string
}

// Extract Wesite Content
func Extract(url string) (sanitizedContent *Content, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	readableDoc, err := readability.NewDocument(string(body))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	readableContent := readableDoc.Content()
	doc, err := goquery.NewDocumentFromResponse(resp)
	title := doc.Find("title").Text()
	return &Content{
		url:              url,
		title:            title,
		readableContent:  readableContent,
		sanitizedContent: sanitize.HTML(readableContent),
	}, nil
}
