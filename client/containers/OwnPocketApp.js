import React, {
  Component
} from 'react';
import { connect } from 'react-redux';
import SearchBookmark from '../components/SearchBookmark';
import AddBookmark from '../components/AddBookmark';
import BookmarkList from '../components/BookmarkList';
import * as actions from '../actions/actions';

export default class OwnPocketApp extends Component {
  render() {
    const {
      dispatch, bookmarks
    } = this.props;
    return (
      < div >
        <SearchBookmark onSearchClick = {
          searchPhrase =>
            dispatch(actions.searchBookmark(searchPhrase))
        }/>
        <AddBookmark onAddClick = {
          url =>
            dispatch(actions.addBookmark(url))
        }/>
        <BookmarkList bookmarks = {
          bookmarks
        }
        onDeleteBookmarkClick = {
          id => dispatch(actions.deleteBookmark(id))
        }/>
      </div>
    );
  }
}

  // Which props do we want to inject, given the global state?
  // Note: use https://github.com/faassen/reselect for better performance.
  function select(state) {
    console.log(state);
    return {
      bookmarks: state.ownPocket
    };
  }

  // Wrap the component to inject dispatch and state into it
  export default connect(select)(OwnPocketApp);
