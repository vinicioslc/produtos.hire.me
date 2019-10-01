import React, { Fragment, Component } from "react";
import Planos from "../../Planos/Planos";
import "./Index.css";

export default class Index extends Component {
  constructor(props) {
    super(props);
    this.state = { theme: "vivo" };
  }

  render() {
    return (
      <div class="App">
        <Planos theme={this.state.theme}></Planos>
      </div>
    );
  }
}
