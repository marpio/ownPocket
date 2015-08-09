package models

// Bookmark db struct
type Bookmark struct {
	Docid             int    `json:"docid"`
	URL               string `json:"url"`
	Title             string `json:"title"`
	ReadableContent   string
	SearchableContent string
}

// SearchBookmarks full text bookmarks
func (db *DB) SearchBookmarks(searchPhrase string) ([]*Bookmark, error) {
	rows, err := db.Query("SELECT docid, url, title FROM bookmarks WHERE readableContent MATCH ?;", searchPhrase)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bks []*Bookmark
	for rows.Next() {
		bk := new(Bookmark)
		err := rows.Scan(&bk.Docid, &bk.URL, &bk.Title)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

// AddBookmark adds bookamr
func (db *DB) AddBookmark(bookmark *Bookmark) (*Bookmark, error) {
	stmt, err := db.Prepare("INSERT INTO bookmarks(url, title, readableContent, sanitizedContent) VALUES (?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(bookmark.URL, bookmark.Title, bookmark.ReadableContent, bookmark.SearchableContent)
	if err != nil {
		return nil, err
	}
	return bookmark, nil
}
