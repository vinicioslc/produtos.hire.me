import React, { Fragment } from "react";

import ThemmedComponent from "../../Base/ThemmedComponent/index";
import DetailsModal from "../DetailsModal/DetailsModal";
import ActionBtn from "../ActionBtn/ActionBtn";
import Price from "../Price/Price";
import "./PlanCard.css";

export default class Plans extends ThemmedComponent {
  curPlan() {
    return this.props.curPlan || {};
  }
  styleClass() {
    return (
      (this.props.styleclass || "") +
      (this.props.isPair ? " dark-bg " : " bright-bg ") +
      this.getTheme()
    );
  }
  curAdvantages() {
    return this.curPlan().plan_advantages || [];
  }

  render() {
    return (
      <li className={`plan-li ${this.getTheme()}`} key={this.curPlan().id}>
        <div className={`plan-container ${this.getTheme()}`}>
          <div className={`plan-header ${this.styleClass()}`}>
            <div className={`plan-title-container ${this.getTheme()}`}>
              <span className={`plan-title-text ${this.getTheme()}`}>
                {this.curPlan().plan_title}
              </span>
            </div>
            <div className={`plan-details-container ${this.getTheme()}`}>
              <span className={`plan-details-text ${this.getTheme()}`}>
                {this.curPlan().plan_details}
              </span>
            </div>
          </div>
          <Highlights
            highlights={this.curPlan()}
            className={this.styleClass()}
          ></Highlights>
          <PlanAdvantages advantages={this.curAdvantages()}></PlanAdvantages>
          <div className={`plan-price-section ${this.getTheme()}`}>
            <PlanFinalSection
              currentPlan={this.curPlan()}
              className={this.getTheme()}
            ></PlanFinalSection>
            <DetailsModal
              button={
                <div className="mt">
                  <ActionBtn theme={this.getTheme()}>Eu Quero !</ActionBtn>
                </div>
              }
            ></DetailsModal>
          </div>
          <SeeDetailsComponent
            theme={this.getTheme()}
            plan={this.curPlan()}
          ></SeeDetailsComponent>
        </div>
      </li>
    );
  }
}

/**
 * Show Plan in featured advantages
 */
const Highlights = function({ highlights, className }) {
  return highlights.plan_highlights.length > 0 ? (
    <div className={`plan-highlight ${className}`}>
      <img
        alt=""
        className="highlights-icon"
        src={highlights.plan_highlights[0].h_icon}
      ></img>
      <div style={{ display: "inline" }}>
        <span className="bold">{highlights.plan_highlights[0].h_title}</span>
        <span>{highlights.plan_highlights[0].h_description}</span>
      </div>
    </div>
  ) : (
    <Fragment></Fragment>
  );
};

/**
 * Show each plan advantages
 */
const PlanAdvantages = function({ advantages, className }) {
  return (
    <ol className="plan-advantages">
      {advantages.map(advantage => {
        return (
          <li key={advantage.a_title} className="advantage">
            <p>
              <span className={`bold ${className}`}>{advantage.a_title}</span>
              <span>{advantage.a_description}</span>
            </p>
          </li>
        );
      })}
    </ol>
  );
};

/**
 * Show Price and som final details
 */
const PlanFinalSection = function({ currentPlan, className }) {
  return (
    <div className={`price-details-section ${className}`}>
      <Price
        price={currentPlan.plan_price}
        days={currentPlan.plan_limit_days}
        theme={className}
      />

      <span className={`plan-small-words ${className}`}>
        {`Válido por ${currentPlan.plan_limit_days || 0} dias`} dias
      </span>
      <span className={`plan-small-words grey ${className}`}>
        {currentPlan.plan_small_words}
      </span>
    </div>
  );
};

const SeeDetailsComponent = function({ plan, theme }) {
  return (
    <div className="plan-tray">
      <DetailsModal
        title="Detalhes do plano"
        theme={theme}
        button={
          <ActionBtn className="details" theme={theme}>
            Ver Detalhes
          </ActionBtn>
        }
      >
        <div
          dangerouslySetInnerHTML={{
            __html: plan.plan_more_details.franchise
          }}
        ></div>
      </DetailsModal>
    </div>
  );
};
