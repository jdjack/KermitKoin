import React, { Component } from 'react';
import axios from 'axios';
import Button from './Button';
import firebase from './firebase'
import sha256 from 'sha256';

class RegisterForm extends Component {

  constructor(props) {
    super(props);
    this.handleRegisterSuccess = this.handleRegisterSuccess.bind(this);

    // Handle when the form updates
    this.handleUsernameChange = this.handleUsernameChange.bind(this);
    this.handlePasswordChange = this.handlePasswordChange.bind(this);
    this.handlePasswordAgainChange = this.handlePasswordAgainChange.bind(this);

    // Connect to the firebase database
    this.state = {
      'username': "",
      'password': "",
      'password_again': "",
      'show_password_error': false,
      'user_exists_error': "",
      'wallet_address': ""
    }

    this.handleSubmit = this.handleSubmit.bind(this);

  }

  handleUsernameChange(event) {
    this.setState({ username: event.target.value });
  }
  handlePasswordChange(event) {
    this.setState({ password: event.target.value });
  }
  handlePasswordAgainChange(event) {
    this.setState({ password_again: event.target.value });
  }

  handleSubmit(event) {

    event.preventDefault();

    // Check that the password and password_again are the same
    if (this.state.password !== this.state.password_again) {
      this.setState({"show_password_error": true});
      return;
    } else {
      if (this.state.show_password_error) {
        this.setState({"show_password_error": false});
      }
    }

    // The address of the host IP
    const host_ip = "129.31.236.46";

    // TODO: Check not already registered

    // Get the number of seconds (the unix timestamp)
    const dateTime = Date.now();
    const timestamp = Math.floor(dateTime / 1000);

    console.log(timestamp);

    // Send request to wallet to get the user's wallet address
    // TODO: tweak parameters
    const req = 'http://' + host_ip + ':8082/getAddress?key=' + timestamp;
    console.log(req);
    axios.get('http://' + host_ip + ':8082/getAddress?key=' + timestamp)
    .then(response => this.setState({wallet_addr: response.data.name}),
          (error) => { console.log(error) })
    .then(function(result) {

      // Now add this user to the database
      const usersRef = firebase.database().ref('users');

      const user = {
        username: this.state.username,
        password: sha256(this.state.password),
        timestamp: timestamp,
        wallet_addr: this.state.wallet_addr
      }

      usersRef.push(user);

    });
    // TODO: update on response

  }

  handleRegisterSuccess(event) {
    this.props.handleRegisterSuccess();
  }

  render() {

    const show_password_error = this.state.show_password_error;
    var password_error = null;
    if (show_password_error) {
      password_error = <p className="ErrorMsg">
                         The two passwords did not match.
                      </p>;
    }

    var user_exists_error = null;
    if (this.state.user_exists_error != "") {
      console.log(this.state.user_exists_error);
      user_exists_error = <p className="ErrorMsg">
                            The username <i>{this.state.user_exists_error}</i> already exists.
                          </p>;
    }

    return (
      <div id="RegisterPage">
        <h4>REGISTERING</h4>

        <div id="RegisterForm">
          <form onSubmit={this.handleSubmit}>
            <label>
              Username:
              <input type="text" value={this.state.value} onChange={this.handleUsernameChange} />
            </label>
            {user_exists_error}
            <label>
              Password:
              <input type="password" value={this.state.value} onChange={this.handlePasswordChange} />
            </label>
            <label>
              Password Confirmation:
              <input type="password" value={this.state.value} onChange={this.handlePasswordAgainChange} />
            </label>
            {password_error}
            <input type="submit" value="Submit" />
          </form>
        </div>

        <Button
          onClick={this.handleRegisterSuccess}
          name={"Back"}
          />
      </div>
    );
  }

}

export default RegisterForm
