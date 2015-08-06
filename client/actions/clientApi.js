export function searchBookmarks(searchPhrase) {
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
