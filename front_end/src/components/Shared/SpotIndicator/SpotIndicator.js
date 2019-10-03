import React, { Fragment } from "react";
import ThemmedComponent from "../../Base/ThemmedComponent";
import "./SpotIndicator.css";

export default class EmptyResults extends ThemmedComponent {
  constructor(props) {
    super(props);
  }

  showSpot() {
    return this.props.showSpot || false;
  }
  render() {
    return this.showSpot() ? (
      <div className={`spot-container ${this.getTheme()}`}>
        <div className={`spot-indicator ${this.getTheme()}`}>
          <span className={`spot-span ${this.getTheme()}`}></span>
        </div>
      </div>
    ) : (
      <Fragment></Fragment>
    );
  }
}
