import React, { Component } from 'react';
import Button from './Button';

class LoginForm extends Component {

  constructor(props) {

    super(props);
    this.state = {username: "", password: ""}

    this.unameHandleChange = this.unameHandleChange.bind(this);
    this.pwordHandleChange = this.pwordHandleChange.bind(this);

    this.handleRegisterClick = this.handleRegisterClick.bind(this);

    this.handleSubmit = this.handleSubmit.bind(this);

  }

  unameHandleChange(event) {
    this.setState({username: event.target.value});
  }
  pwordHandleChange(event) {
    this.setState({password: event.target.value});
  }

  handleRegisterClick(event) {
    this.props.handleRegisterClick();
  }

  handleSubmit(event) {
    // TODO: authenticate with Firebase
    event.preventDefault();
    // Clear the current state
    this.setState({username: "", password: ""});
    var logged_in_success = true;
    if (logged_in_success) {
      this.props.handleLoggedInSuccess();
    }
  }

  render() {
    return (
      <div id="LoginForm">
        <form onSubmit={this.handleSubmit}>
          <label>
            Username:
            <input type="text" value={this.state.value} onChange={this.unameHandleChange} />
          </label>
          <label>
            Password:
            <input type="password" value={this.state.value} onChange={this.pwordHandleChange} />
          </label>
          <input type="submit" value="Submit" />
        </form>
        <div id="RegisterButton">
          <Button
            onClick={this.handleRegisterClick}
            name={"Register"} />
        </div>
      </div>
    );
  }

}

export default LoginForm
