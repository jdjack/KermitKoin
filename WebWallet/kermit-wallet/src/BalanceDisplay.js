import React, { Component } from 'react';
import axios from 'axios';

class BalanceDisplay extends Component {

  constructor(props) {
    super(props);

    this.state = {
      'balance': 0.0
    }

    this.getWalletID = this.getWalletID.bind(this);

    this.getBalance = this.getBalance.bind(this);

    console.log(this.props);

    const host_ip = "129.31.236.46";

    axios.get('http://' + host_ip + ':8080/getBalance?id=' + this.props.getWalletID())
    .then(response => {
      console.log("RESPONSE");
      console.log(response);
      this.setState({ balance: response.data.balance})
    },
    (error) => { console.log(error) });

  }

  getBalance(event) {
    return this.state.balance;
  }

  getWalletID(event) {
    return this.props.getWalletID();
  }

  render() {
    return(
      <div id="BalanceScreen">
        <p id="BalanceDisplayContent">
        {this.getBalance()}
        </p>
      </div>
    );
  }

}

export default BalanceDisplay
