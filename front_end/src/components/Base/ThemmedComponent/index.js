import { Component } from "react";

class MissingThemeError extends Error {
  message = "MissingThemeError";
  constructor(args) {
    super(args);
    this.stack = Error().stack();
  }
}

export default class ThemmedComponent extends Component {
  constructor(props) {
    super(props);
    if (!this.props.theme) throw new MissingThemeError();
  }
  getTheme() {
    return this.props.theme || "";
  }
}
