import React, { Fragment, Component } from "react";

import plansApi from "../../services/plans-api";
import PlanCard from "../Shared/PlanCard/PlanCard";
import "./Planos.css";
import EmptyResults from "../EmptyResults/EmptyResults";
export default class Planos extends Component {
  constructor(props) {
    super(props);
    this.state = {};
    plansApi.getPlans(this.props.theme).then(plans => {
      console.log("FETCHED ", plans);
      this.setState({
        plans: plans
      });
    });
  }

  render() {
    return (
      <Fragment>
        <div className={`app ${this.props.theme}`}>
          <header className={`header ${this.props.theme}`}>
            <a className={`logo ${this.props.theme}`}></a>
          </header>
          <h1 className={`offer-title ${this.props.theme}`}>
            Ofertas feitas para você
          </h1>
          {this.state.plans ? (
            <ul>
              <PlansList
                plans={this.state.plans}
                customClass={this.props.customClass}
                theme={this.props.theme}
              ></PlansList>
            </ul>
          ) : (
            <EmptyResults></EmptyResults>
          )}
        </div>
      </Fragment>
    );
  }
}

const PlansList = ({ plans, customClass, theme }) => {
  return plans.map((curPlan, index) => {
    let isPair = (index & 1) == 0;
    return (
      <PlanCard
        curPlan={curPlan}
        isPair={isPair}
        className={customClass}
        theme={theme}
      ></PlanCard>
    );
  });
};
