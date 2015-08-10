import React, { Component, PropTypes } from 'react';
import Bookmark from './Bookmark';

export default class BookmarkList extends Component {
  render() {
    return (
      <ul>
        {this.props.bookmarks.map((bookmark, index) =>
          <Bookmark {...bookmark}
                key={bookmark.docid}
                handleDeleteClick={() => {
                  this.props.onDeleteBookmarkClick(bookmark.docid);
                }}/>
        )}
      </ul>
    );
  }
}

BookmarkList.propTypes = {
  bookmarks: PropTypes.arrayOf(PropTypes.shape({
    url: PropTypes.string.isRequired
  }).isRequired).isRequired
};
