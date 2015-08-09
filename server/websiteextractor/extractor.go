package websiteextractor

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kennygrant/sanitize"
	"github.com/marpio/ownPocket/server/models"
	"github.com/mauidude/go-readability"
	"io/ioutil"
	"log"
	"net/http"
)

func extractTitle(url string, c chan string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Println(err)
	}
	c <- doc.Find("title").First().Text()
}

func extractReadableContent(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	readableDoc, err := readability.NewDocument(string(body))
	if err != nil {
		log.Fatal(err)
	}
	c <- readableDoc.Content()
}

// Extract Wesite Content
func Extract(url string) (bookmark *models.Bookmark, err error) {
	c := make(chan string)
	go extractTitle(url, c)
	go extractReadableContent(url, c)
	title, readableContent := <-c, <-c

	log.Println(title)

	return &models.Bookmark{
		URL:               url,
		Title:             title,
		ReadableContent:   readableContent,
		SearchableContent: sanitize.HTML(readableContent),
	}, nil
}
