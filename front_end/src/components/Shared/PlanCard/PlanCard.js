import React, { Fragment, Component } from "react";

import DetailsModal from "../../DetailsModal/DetailsModal";
import ActionBtn from "../ActionBtn/ActionBtn";
import Price from "../Price/Price";

export default class Planos extends Component {
  constructor(props) {
    super(props);
  }
  curPlan() {
    return this.props.curPlan || {};
  }
  styleClass() {
    return (
      (this.props.styleclass || "") +
      (this.props.isPair ? " dark-bg " : " bright-bg ") +
      this.props.theme
    );
  }
  render() {
    return (
      <li className="plan-li" key={this.curPlan().id}>
        <div className="plan-container">
          <div className={`plan-header ${this.styleClass()}`}>
            <span className="plan-title">{this.curPlan().plan_title}</span>
            <span className="plan-title-details">
              {this.curPlan().plan_details}
            </span>
          </div>
          {this.curPlan().plan_highlights.length > 0 ? (
            <div className={`plan-highlight ${this.styleClass()}`}>
              <img
                className="highlights-icon"
                src={this.curPlan().plan_highlights[0].h_icon}
              ></img>
              <div style={{ display: "inline" }}>
                <span className="bold">
                  {this.curPlan().plan_highlights[0].h_title}
                </span>
                <span>{this.curPlan().plan_highlights[0].h_description}</span>
              </div>
            </div>
          ) : (
            <Fragment></Fragment>
          )}
          <ol className="plan-advantages">
            {this.curPlan().plan_advantages.map(advantage => {
              return (
                <li className="advantage">
                  <p>
                    <span className="bold">{advantage.a_title}</span>
                    <span>{advantage.a_description}</span>
                  </p>
                </li>
              );
            })}
          </ol>
          <div className="plan-end-section">
            <Price
              price={this.curPlan().plan_price}
              days={this.curPlan().plan_limit_days}
              theme={this.props.theme}
            />

            <span class="plan-small-words" style={{ display: "block" }}>
              {`Válido por ${this.curPlan().plan_limit_days || 0} dias`} dias
            </span>
            <p class="plan-small-words">{this.curPlan().plan_small_words}</p>
          </div>
          <DetailsModal
            button={
              <div className="mt">
                <ActionBtn theme={this.props.theme}>Assinar</ActionBtn>
              </div>
            }
          ></DetailsModal>
          <div className="plan-tray">
            <DetailsModal
              title="Detalhes do plano"
              theme={this.props.theme}
              button={
                <ActionBtn className="details" theme={this.props.theme}>
                  Ver Detalhes
                </ActionBtn>
              }
            >
              <div
                dangerouslySetInnerHTML={{
                  __html: this.curPlan().plan_more_details.franchise
                }}
              ></div>
            </DetailsModal>
          </div>
        </div>
      </li>
    );
  }
}
