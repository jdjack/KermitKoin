import React, { Component } from 'react';
import Button from './Button';
import firebase from './firebase';
import sha256 from 'sha256';

class LoginForm extends Component {

  constructor(props) {

    super(props);
    this.state = {username: "",
                  password: "",
                  firebase_password: "",
                  wallet_id: "",
                  auth_user: false
                }

    this.unameHandleChange = this.unameHandleChange.bind(this);
    this.pwordHandleChange = this.pwordHandleChange.bind(this);

    this.handleRegisterClick = this.handleRegisterClick.bind(this);

    this.handleSubmit = this.handleSubmit.bind(this);

    this.setWalletID = this.setWalletID.bind(this);

  }

  setWalletID(event) {
    this.props.setWalletID(this.state.wallet_id);
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
    const usersRef = firebase.database().ref('users');

    var that = this;
    usersRef.orderByChild('username').equalTo(this.state.username).on("value", function(snapshot) {
        console.log(snapshot.val());
        var res = snapshot.val();
        that.setState({ "firebase_password": res[Object.keys(res)[0]]["password"]});
        that.setState({ "wallet_id": res[Object.keys(res)[0]]["wallet_addr"]});
        that.setWalletID(that.state.wallet_id);
        console.log(that.state);

        // Check that the sha of the password is correct
        if (sha256(that.state.password) == that.state.firebase_password) {
          console.log(that.props);
          that.props.handleLoggedInSuccess();
        }

    });

    // var that = this;
    // usersRef.orderByChild('username').equalTo(this.state.username).then(function(snapshot) {
    //   that.setState({
    //     firebase_password: snapshot.val()['password']
    //   });
    // });



    event.preventDefault();
    // Clear the current state
    // this.setState({username: "", password: ""});
    // var logged_in_success = true;
    // if (logged_in_success) {
    //   this.props.handleLoggedInSuccess();
    // }
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
