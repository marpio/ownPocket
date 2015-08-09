import {
  ADD_BOOKMARK, DELETE_BOOKMARK, SEARCH_BOOKMARK
}
from '../constants/ActionTypes';
/*
Docid             int64  `json:"docid"`
URL               string `json:"url"`
Title             string `json:"title"`
FirstParagraph    string `json:"firstParagraph"`


*/
const initialState = [{
  docid: 111,
  url: "some.url",
  title: "title1",
  firstParagraph: "firstParagraph"
}];

export default function bookmarks(state = initialState, action) {
  switch (action.type) {
    case ADD_BOOKMARK:
      return [action.bookmark, ...state];
    case DELETE_BOOKMARK:
      return state.filter(b => b.docid !== action.docid);
    case SEARCH_BOOKMARK:
      if(action.searchPhrase === "") return state;
      return action.searchedBookmarks;
    default:
      return state;
  }
}
