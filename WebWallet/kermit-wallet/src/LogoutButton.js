import React, { Component } from 'react';

class LogoutButton extends Component {

  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    this.props.handleLogoutClick();
  }

  render() {
    return (
      <button onClick={this.handleClick} id="LogOutButton">
        Log Out
      </button>
    );
  }

}

export default LogoutButton
