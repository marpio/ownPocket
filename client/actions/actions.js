import * as types from '../constants/ActionTypes';
import * as clientApi from './clientApi';

export function addBookmark(url) {
  return function(dispatch) {
    return clientApi.addBookmark(url).then(
      addedBookmark => dispatch({
        type: types.ADD_BOOKMARK,
        bookmark: addedBookmark
      })
    )
  };
}

export function deleteBookmark(id) {
  return {
    type: types.DELETE_BOOKMARK,
    id
  }
}

export function searchBookmark(searchPhrase) {
  return function(dispatch) {
    return clientApi.searchBookmarks(searchPhrase).then(
      searchedBookmarks => dispatch({
        type: types.SEARCH_BOOKMARK,
        searchPhrase,
        searchedBookmarks
      })
    )
  };
}
