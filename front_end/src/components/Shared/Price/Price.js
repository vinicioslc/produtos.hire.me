import React, { Fragment, Component } from "react";
import "./Price.css";
import ThemmedComponent from "../../Base/ThemmedComponent";
export default class Price extends ThemmedComponent {
  priceSymbol() {
    return navigator.language === "pt-BR" ? "R$" : "$";
  }
  priceUnit() {
    return this.props.price.toString().split(".")[0];
  }
  priceCent() {
    return this.props.price.toString().split(".")[1];
  }

  render() {
    return (
      <Fragment>
        <div className={`price-box offer-text bold ${this.getTheme()}`}>
          <span className={`offer-price-cifr ${this.getTheme()}`}>
            {this.priceSymbol()}
          </span>
          <span className={`offer-price-int ${this.getTheme()}`}>
            {this.priceUnit()}
          </span>
          <span className={`offer-price-cent ${this.getTheme()}`}>
            <span className={`offer-price-cent-split ${this.getTheme()}`}>
              ,
            </span>
            {this.priceCent()}
          </span>
        </div>
      </Fragment>
    );
  }
}
