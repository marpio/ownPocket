package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marpio/ownPocket/server/dto"
	"github.com/marpio/ownPocket/server/models"
	"github.com/marpio/ownPocket/server/websiteextractor"
	"github.com/rs/cors"
)

type handler func(w http.ResponseWriter, r *http.Request)

func basicAuth(pass handler) handler {

	return func(w http.ResponseWriter, r *http.Request) {

		auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "bad syntax", http.StatusBadRequest)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		pass(w, r)
	}
}

func validate(username, password string) bool {
	if username == "username" && password == "password" {
		return true
	}
	return false
}

// Env struct
type Env struct {
	db models.Datastore
}

func checkErrorAndWriteToRes(err error, w http.ResponseWriter) (shouldReturn bool) {
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

// AddBookmarkHandler request handler
func (env *Env) AddBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var url dto.URL
	err := decoder.Decode(&url)
	if checkErrorAndWriteToRes(err, w) {
		return
	}
	c, err := websiteextractor.Extract(url.URL)

	bookmark, err := env.db.AddBookmark(c)
	log.Println(bookmark.Docid)
	if checkErrorAndWriteToRes(err, w) {
		return
	}
	b, err := json.Marshal(bookmark)
	if checkErrorAndWriteToRes(err, w) {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

// SearchBookmarkHandler request handler
func (env *Env) SearchBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	bookmarks, err := env.db.SearchBookmarks(q)
	if checkErrorAndWriteToRes(err, w) {
		return
	}
	bs, err := json.Marshal(bookmarks)
	if checkErrorAndWriteToRes(err, w) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := ":8089"
	db, err := models.NewDB("./ownPocket.db")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}
	log.Printf("serving on port %s\n", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/bookmarks/create", basicAuth(env.AddBookmarkHandler))
	mux.HandleFunc("/api/bookmarks/search", basicAuth(env.SearchBookmarkHandler))
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(port, handler)

	// http.HandleFunc("/api/bookmarks/create", env.AddBookmarkHandler)
	// http.HandleFunc("/api/bookmarks/search", env.SearchBookmarkHandler)
	// http.ListenAndServe(port, nil)
}
