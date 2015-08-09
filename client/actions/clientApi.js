import request from "superagent";

export function searchBookmarks(searchPhrase) {
  console.log(request);
  request.get('http://localhost:8089/api/bookmarks/search?q=tutorial', function(err, res){
  if (err) throw err;
  console.log(res.text);
});


  return new Promise(function(resolve, reject) {
    resolve([{
      id: 0,
      title: "title0",
      tags: ["tag0"]
    }, {
      id: 1,
      title: "title1",
      tags: ["tag1"]
    }, {
      id: 2,
      title: "title2",
      tags: ["tag2"]
    }, {
      id: 3,
      title: "title3",
      tags: ["tag3"]
    }].filter(b => b.title === searchPhrase));
  });
}

export function addBookmark(url) {
  return new Promise(function(resolve, reject) {
    resolve({
      id: 10,
      title: "title10",
      tags: ["tag10"]
    })
  });
}
