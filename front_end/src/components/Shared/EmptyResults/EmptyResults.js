import React from "react";
import ThemmedComponent from "../../Base/ThemmedComponent";

export default class EmptyResults extends ThemmedComponent {
  render() {
    return (
      <div>
        <h2
          style={{ color: "grey", fontSize: "5em" }}
          className={`${this.getTheme()}`}
        >
          {"Sem planos para você :'("}
        </h2>
      </div>
    );
  }
}
