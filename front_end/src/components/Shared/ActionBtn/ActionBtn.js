import React, { Fragment, Component } from "react";
import "./ActionBtn.css";
export default class ActionBtn extends Component {
  constructor(props) {
    super(props);
  }
  className() {
    return this.props.className || " ";
  }
  render() {
    return (
      <Fragment>
        <button
          className={`btn ${this.className()} ${this.props.theme || "yellow"} `}
        >
          {this.props.children}
        </button>
      </Fragment>
    );
  }
}
