import React, { Component } from 'react';
import LogoutButton from './LogoutButton'
import BalanceDisplay from './BalanceDisplay'
import SendDisplay from './SendDisplay'
import RecieveDisplay from './RecieveDisplay'
import Button from './Button'

class UserDashboard extends Component {

  constructor(props) {
    super(props);
    this.handleLogoutClick = this.handleLogoutClick.bind(this);
    this.handleRecieveClick = this.handleRecieveClick.bind(this);
    this.handleSendClick = this.handleSendClick.bind(this);
    this.handleBalanceClick = this.handleBalanceClick.bind(this);
    this.getWalletID = this.getWalletID.bind(this);
    this.state = {display: "balance"};
  }

  getWalletID(event) {
    return this.props.getWalletID();
  }

  handleLogoutClick(event) {
    this.props.handleLogoutClick();
  }

  handleBalanceClick(event) {
    this.setState({display: "balance"});
  }

  handleRecieveClick(event) {
    this.setState({display: "recieve"});
  }

  handleSendClick(event) {
    this.setState({display: "send"});
  }

  render() {

    // Which display to render
    const display = this.state.display;

    let content = null;
    switch (display) {
      case "send":
        content = <SendDisplay
                    getWalletID={this.getWalletID} />;
        break;
     case "recieve":
        content = <RecieveDisplay
                    getWalletID={this.getWalletID} />;
        break;
      case "balance":
      default:
        content = <BalanceDisplay
                     getWalletID={this.getWalletID} />;
    }

    let buttons = null;
    switch(display) {
      case "send":
      case "recieve":
        buttons =
          <div id="DashboardButtons">
            <Button
              onClick={this.handleBalanceClick}
              name={"Back"} />
          </div>;
        break;
      case "balance":
      default:
        buttons =
          <div id="DashboardButtons">
            <Button
              onClick={this.handleSendClick}
              name={"Send"} />
            <Button
              onClick={this.handleRecieveClick}
              name={"Recieve"} />
          </div>;
        break;
    }

    return (
      <div id="UserDashboard">
        {content}
        {buttons}
        <LogoutButton
        handleLogoutClick={this.handleLogoutClick} />
      </div>
    );
  }

}

export default UserDashboard
