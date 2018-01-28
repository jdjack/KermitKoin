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

    this.state = {isLoggedIn: false};
    this.state = {isRegistering: false};
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
                  handleRegisterClick={this.handleRegisterClick} />;
    } else {
      content = <UserDashboard
                  handleLogoutClick={this.handleLogoutClick} />;
    }

    return (
      <div>
        <h2>Welcome!</h2>
        {content}
      </div>
    );
  }

}


export default LoginControl;
