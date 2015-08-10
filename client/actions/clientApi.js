import request from "superagent";
import sp from "superagent-promise";

let agent = sp(request, Promise);
export function searchBookmarks(searchPhrase) {

  return agent.get("http://localhost:8089/api/bookmarks/search?q=" + searchPhrase).end();

  // t.then(r => console.log(r));
  //
  // return new Promise(function(resolve, reject) {
  //   resolve([{
  //     id: 0,
  //     title: "title0",
  //     tags: ["tag0"]
  //   }, {
  //     id: 1,
  //     title: "title1",
  //     tags: ["tag1"]
  //   }, {
  //     id: 2,
  //     title: "title2",
  //     tags: ["tag2"]
  //   }, {
  //     id: 3,
  //     title: "title3",
  //     tags: ["tag3"]
  //   }].filter(b => b.title === searchPhrase));
  // });
}

export function addBookmark(url) {
  return agent.post("http://localhost:8089/api/bookmarks/create")
              .send({ url: url })
              .end();
  // return new Promise(function(resolve, reject) {
  //   resolve({
  //     id: 10,
  //     title: "title10",
  //     tags: ["tag10"]
  //   })
  // });
}
