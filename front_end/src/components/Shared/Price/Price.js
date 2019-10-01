import React, { Fragment, Component } from "react";
import "./Price.css";
export default class Price extends Component {
  constructor(props) {
    super(props);
  }
  priceSymbol() {
    return navigator.language == "pt-BR" ? "R$" : "$";
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
        <div
          className={`price-box offer-text primary-text-color bold ${this.props
            .theme || ""}`}
        >
          <span className="offer-price-cifr">{this.priceSymbol()}</span>
          <span className="offer-price-int">{this.priceUnit()}</span>
          <span className="offer-price-cent">
            <span className="offer-price-cent-split">,</span>
            {this.priceCent()}
          </span>
        </div>
      </Fragment>
    );
  }
}
