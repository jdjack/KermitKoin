import React, { Component } from 'react';
import axios from 'axios';

class SendDisplay extends Component {

  constructor(props) {
    super(props);

    this.state = {
      'balance': 0.0,
      'wallet_addr_reciever': "",
      'amount_to_send': 0.0,
      'success_send': false,
      'attempted_sent': false
    }


    this.handleDestAddrChange = this.handleDestAddrChange.bind(this);
    this.handleAmountChange = this.handleAmountChange.bind(this);

    this.getWalletID = this.getWalletID.bind(this);

    const host_ip = "129.31.236.46";

    axios.get('http://' + host_ip + ':8080/getBalance?id=' + this.props.getWalletID())
    .then(response => {
      this.setState({ balance: response.data.balance})
    },
    (error) => { console.log(error) });
  }

  getWalletID(event) {
    return this.props.getWalletID();
  }

  handleDestAddrChange(event) {
    this.setState({ 'wallet_addr_reciever': event.target.value});
  }

  handleAmountChange(event) {
    this.setState({ 'amount_to_send': event.target.value});
  }


  handleSubmit(event) {

    event.preventDefault();

    // The address of the host IP
    const host_ip = "129.31.236.46";

    // Send request to wallet to get the user's wallet address
    // TODO: tweak parameters
    axios.get('http://' + host_ip + ':8080/makeTransaction?ownID= ' + this.getWalletID() + '&destID=' + this.state.wallet_addr_reciever + '&amount=' + this.state.amount_to_send)
    .then(response => {
      this.setState({ 'success_send': true });
      console.log(response);
    },
      (error) => { this.setState({ 'success_send': false }); console.log(error) });

  }

  render() {

    // Whether to display the success message or not
    var form = null;
    var success_message = null;
    var error_message = null;
    if (!this.state.attempted_sent || !this.state.success_send) {
      form =
          <div id="SendMoneyForm">
          <div id="CurrentBalanceSend">
                <p>Available balance: {this.state.balance}</p>
              </div>

              <div id="SendForm">
                <form onSubmit={this.handleSubmit}>
                  <label>
                    Address to send to:
                    <input type="text" value={this.state.value} placeholder="0x1234567890ABCDEF" onChange={this.handleDestAddrChange} />
                  </label>
                  <label>
                    Amount to send:
                    <input type="test" value={this.state.value} placeholder="1234" onChange={this.handleAmountChange} />
                  </label>
                  <input type="submit" value="Send" />
                </form>
              </div>
            </div>;

    }
    if (!this.state.success_send && this.state.attempted_sent) {
      error_message =
        <p class="ErrorMsg">Failed to send. Please try again.</p>;
    }
    if (this.state.success_send) {
      success_message =
      <p class="SendMsg">Sent successfully!</p>;
    }

    return(
      <div id="SendPage">
        {form}
        {error_message}
        {success_message}
      </div>
    );
  }

}

export default SendDisplay
