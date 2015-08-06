import React, { findDOMNode, Component, PropTypes } from 'react';

export default class SearchBookmark extends Component {
  render() {
    return (
      <div>
        <input type='text' ref='input' />
        <button onClick={(e) => this.handleClick(e)}>
          Search
        </button>
      </div>
    );
  }

  handleClick(e) {
    const node = findDOMNode(this.refs.input);
    const searchText = node.value.trim();
    this.props.onSearchClick(searchText);
    // node.value = '';
  }
}

SearchBookmark.propTypes = {
  onSearchClick: PropTypes.func.isRequired
};
