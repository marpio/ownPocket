export function searchBookmarks(searchPhrase) {
   return window.fetch("http://localhost:8089/api/bookmarks/search?q=" + searchPhrase, {method: "get"}).then(function(response) {
	// Convert to JSON
	return response.json();
});
}

export function addBookmark(url) {
  return window.fetch("http://localhost:8089/api/bookmarks/create", {
	method: 'post',
	body: JSON.stringify({ url: url })
}).then(function(response) {
return response.json();
});
}
