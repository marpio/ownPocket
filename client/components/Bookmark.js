import React, { Component, PropTypes } from 'react';

export default class Bookmark extends Component {
  render() {
    return (
      <li
        style={{
          textDecoration: "none",
          cursor: "default"
        }}>
        {this.props.title} <button onClick={this.props.handleDeleteClick}>Delete</button>
      </li>
    );
  }
}

Bookmark.propTypes = {
  handleDeleteClick: PropTypes.func.isRequired,
  url: PropTypes.string.isRequired
};
