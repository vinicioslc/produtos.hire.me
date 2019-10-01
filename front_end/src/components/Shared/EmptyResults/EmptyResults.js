import React, { Fragment, Component } from "react";

export default class EmptyResults extends React.Component {
  constructor(args) {
    super(args);
  }
  render() {
    return (
      <div>
        <h2 style={{ color: "grey", fontSize: "5em" }}>
          {"Sem planos para você :'("}
        </h2>
      </div>
    );
  }
}
