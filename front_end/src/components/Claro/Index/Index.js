import React, { Fragment, Component } from "react";
import Plans from "../../Shared/Plans/Plans";
import "./Index.css";

export default class Index extends Component {
  constructor(props) {
    super(props);
    this.state = { theme: "claro" };
  }

  render() {
    return (
      <div className="App">
        <Plans theme={this.state.theme}></Plans>
      </div>
    );
  }
}
