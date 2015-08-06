package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marpio/ownPocket/server/dto"
	"github.com/marpio/ownPocket/server/websiteextractor"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside req")
	decoder := json.NewDecoder(r.Body)

	var url dto.URL
	err := decoder.Decode(&url)
	if err != nil {
		log.Fatal(err)
	}
	c, err := contentextractor.Extract(url.URL)
	log.Println(c)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8089", nil)
}
