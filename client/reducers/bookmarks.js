import {
  ADD_BOOKMARK, DELETE_BOOKMARK, SEARCH_BOOKMARK
}
from '../constants/ActionTypes';

const initialState = [{
  id: 111,
  title: "action.url",
  tags: ["tag1111"]
}];

export default function bookmarks(state = initialState, action) {
  switch (action.type) {
    case ADD_BOOKMARK:
      return [action.bookmark, ...state];
    case DELETE_BOOKMARK:
      return state.filter(b => b.id !== action.id);
    case SEARCH_BOOKMARK:
      if(action.searchPhrase === "") return state;
      return action.searchedBookmarks;
    default:
      return state;
  }
}
