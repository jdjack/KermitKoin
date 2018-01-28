import React, { Component } from 'react';

class Button extends Component {

  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    this.props.onClick();
  }

  render() {
    return(
      <button onClick={this.handleClick}>
        {this.props.name}
      </button>
    );
  }

}

export default Button
