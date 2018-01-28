import React, { Component } from 'react';

class RecieveDisplay extends Component {

  constructor(props) {
    super(props);

    this.state = {
      'wallet_addr': 0.0
    }

    this.getWalletID = this.getWalletID.bind(this);

  }

  getWalletID(event) {
    return this.props.getWalletID();
  }

  render() {
    return(
      <div id="RecieveDisplay">
        <h4 id="YourAddressLabel">Your Address:</h4>
        <p id="RecieveDisplayWalletAddress">0x
          {this.getWalletID()}
        </p>
      </div>
    );
  }

}

export default RecieveDisplay
