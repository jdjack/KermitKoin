import React, { Component } from 'react';
import LoginForm from './LoginForm';
import UserDashboard from './UserDashboard';
import RegisterForm from './RegisterForm';
import firebase from './firebase.js';

class LoginControl extends Component {
  constructor(props) {
    super(props);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
    this.handleRegisterClick = this.handleRegisterClick.bind(this);

    this.handleLoggedInSuccess = this.handleLoggedInSuccess.bind(this);
    this.handleRegisterSuccess = this.handleRegisterSuccess.bind(this);

    this.setWalletID = this.setWalletID.bind(this);
    this.getWalletID = this.getWalletID.bind(this);

    this.state = {isLoggedIn: false}
    this.state = {isRegistering: false};
    this.state = {walletID: ""};
  }

  setWalletID(id) {
    console.log("SETTING WALLET ID");
    this.setState({walletID: id});
  }

  getWalletID() {
    return this.state.walletID;
  }

  handleLogoutClick() {
    this.setState({isLoggedIn: false});
  }

  handleRegisterClick() {
    this.setState({isRegistering: true});
  }

  handleLoggedInSuccess() {
    this.setState({isLoggedIn: true});
  }

  handleRegisterSuccess() {
    this.setState({isRegistering: false});
  }

  render() {
    const isLoggedIn = this.state.isLoggedIn;
    const isRegistering = this.state.isRegistering;

    let content = null;
    if (isRegistering && !isLoggedIn) {
      content = <RegisterForm
                  handleRegisterSuccess={this.handleRegisterSuccess} />;
    }
    else if (!isLoggedIn) {
      content = <LoginForm
                  handleLoggedInSuccess={this.handleLoggedInSuccess}
                  handleRegisterClick={this.handleRegisterClick}
                  setWalletID={this.setWalletID} />;
    } else {
      content = <UserDashboard
                  handleLogoutClick={this.handleLogoutClick}
                  getWalletID={this.getWalletID} />;
    }

    return (
      <div>
        {content}
      </div>
    );
  }

}


export default LoginControl;
