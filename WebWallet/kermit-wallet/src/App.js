import React, { Component } from 'react';
import LoginControl from './LoginControl';
import logo from './logo.svg';
import './App.css';

class Page extends Component {
  render() {
    return (
      <div className="Page">
        <Header />
        <Content />
      </div>
    );
  }
}

class Header extends Component {
  render() {
    return (
      <div id="header">
        <h1>Kermit Koin</h1>
      </div>
    );
  }
}

class Content extends Component {
  render() {
    return (
      <LoginControl />
    );
  }
}

class App extends Component {
  render() {
    return (
      <Page />
    );
  }
}

export default App;
