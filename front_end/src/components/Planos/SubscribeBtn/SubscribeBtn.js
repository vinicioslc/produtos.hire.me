import React, { Fragment, Component } from "react";
import "./SubscribeBtn.css";
export default class Planos extends Component {
  constructor(props) {
    super(props);
  }
  color() {
    return this.props.color || "yellow";
  }
  render() {
    return (
      <Fragment>
        <button className={`btn ${this.color()}`}>{this.props.children}</button>
      </Fragment>
    );
  }
}
