import React, { Fragment } from "react";

import ThemmedComponent from "../../Base/ThemmedComponent/index";
import DetailsModal from "../DetailsModal/DetailsModal";
import ActionBtn from "../ActionBtn/ActionBtn";
import Price from "../Price/Price";
import "./PlanCard.css";

export default class Plans extends ThemmedComponent {
  getPlan() {
    return this.props.curPlan || {};
  }
  currentStyle() {
    return (
      (this.props.className || "") +
      (this.props.isPair ? " dark-bg " : " bright-bg ") +
      this.getTheme()
    );
  }
  curAdvantages() {
    return this.getPlan().plan_advantages || [];
  }

  render() {
    return (
      <li className={`plan-li ${this.getTheme()}`} key={this.getPlan().id}>
        <div className={`plan-container ${this.getTheme()}`}>
          <PlanHeader
            plan={this.getPlan()}
            theme={this.getTheme()}
            style={this.currentStyle()}
          />
          <Highlights
            highlights={this.getPlan().plan_highlights}
            className={this.currentStyle()}
            theme={this.getTheme()}
          ></Highlights>
          <PlanAdvantages advantages={this.curAdvantages()}></PlanAdvantages>
          <PlanPriceSection
            plan={this.getPlan()}
            theme={this.getTheme()}
          ></PlanPriceSection>
          <SeeDetailsComponent
            theme={this.getTheme()}
            plan={this.getPlan()}
          ></SeeDetailsComponent>
        </div>
      </li>
    );
  }
}

const PlanPriceSection = function({ plan, theme }) {
  return (
    <div className={`plan-price-section ${theme}`}>
      <PlanFinalSection currentPlan={plan} className={theme}></PlanFinalSection>
      <DetailsModal
        button={
          <div className="mt">
            <ActionBtn className={theme}> Quero !</ActionBtn>
          </div>
        }
      ></DetailsModal>
    </div>
  );
};
/**
 * Show Plan in featured advantages
 */
const Highlights = function({ highlights, className, theme }) {
  return highlights.length > 0 ? (
    <div className={`plan-highlight ${className}`}>
      <img alt="" className="highlights-icon" src={highlights[0].h_icon}></img>
      <span className={`highlight-title ${theme}`}>
        {highlights[0].h_title}
      </span>
      <span className={`highlight-desc ${theme}`}>
        {highlights[0].h_description}
      </span>
      <div style={{ display: "inline" }}></div>
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
    <ol className={`plan-advantages`}>
      {advantages.map(advantage => {
        return (
          <li key={advantage.a_title} className={`advantage`}>
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

const PlanHeader = function({ plan, theme, style }) {
  return (
    <div className={`plan-header ${style}`}>
      <div className={`plan-title-container ${theme}`}>
        <span className={`plan-title-text ${theme}`}>{plan.plan_title}</span>
      </div>
      <div className={`plan-details-container ${theme}`}>
        <span className={`plan-details-text ${theme}`}>
          {plan.plan_details}
        </span>
      </div>
    </div>
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
        {`Válido por ${currentPlan.plan_limit_days || 0} ${
          currentPlan.plan_limit_days > 1 ? "dias" : "dia"
        }`}
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
