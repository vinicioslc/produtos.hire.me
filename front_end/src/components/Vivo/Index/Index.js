import React, { Fragment, Component } from "react";
import Planos from "../../Planos/Planos";
import "./Index.css";

export default class Index extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div class="App">
        <div className="vivo-app">
          <header className="vivo-header">
            <a className="vivo-logo"></a>
          </header>
          <Planos></Planos>
        </div>
      </div>
    );
  }
}
