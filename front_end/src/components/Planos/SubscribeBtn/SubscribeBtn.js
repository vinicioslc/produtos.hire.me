import React, { Fragment, Component } from "react";
import "./SubscribeBtn.css";
export default class Planos extends Component {
  constructor(props) {
    super(props);
  }
  render() {
    return (
      <Fragment>
        <button className={`btn yellow`}>Assinar</button>
      </Fragment>
    );
  }
}
