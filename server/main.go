package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marpio/ownPocket/server/dto"
	"github.com/marpio/ownPocket/server/models"
	"github.com/marpio/ownPocket/server/websiteextractor"
)

// Env struct
type Env struct {
	db models.Datastore
}

// AddBookmarkHandler request handler
func (env *Env) AddBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var url dto.URL
	err := decoder.Decode(&url)
	if err != nil {
		log.Println(err)
		http.Error(w, `Could not parse json!`, http.StatusBadRequest)
		return
	}
	c, err := websiteextractor.Extract(url.URL)

	bookmark, err := env.db.AddBookmark(c)
	if err != nil {
		log.Println(err)
		http.Error(w, `Could not save website content!`, http.StatusInternalServerError)
	}
	b, err := json.Marshal(bookmark)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// SearchBookmarkHandler request handler
func (env *Env) SearchBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	bookmarks, err := env.db.SearchBookmarks(q)
	if err != nil {
		log.Println(err)
	}
	bs, err := json.Marshal(bookmarks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db, err := models.NewDB("./ownPocket.db")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}
	log.Println("serving")
	http.HandleFunc("/api/create", env.AddBookmarkHandler)
	http.HandleFunc("/api/search", env.SearchBookmarkHandler)
	http.ListenAndServe(":8089", nil)
}
