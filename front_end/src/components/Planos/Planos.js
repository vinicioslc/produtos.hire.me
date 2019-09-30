import React, { Fragment, Component } from "react";
import Price from "./Price/Price";
import SubscribeBtn from "./SubscribeBtn/SubscribeBtn";
import DetailsModal from "../DetailsModal/DetailsModal";
import plansApi from "../../services/plans-api";
import "./Planos.css";
export default class Planos extends Component {
  constructor(props) {
    super(props);
    this.state = {};
    plansApi.getPlans("vivo").then(plans => {
      console.log("FETCHED ", plans);
      this.setState({
        plans: plans
      });
    });
  }

  render() {
    return (
      <Fragment>
        <h1 className="offer-title">Ofertas feitas para você</h1>
        {this.state.plans ? (
          <ul>
            {this.state.plans.map((curPlan, index) => {
              let isImpar = (index & 1) == 1;
              let backgroundColor = isImpar ? "dark-bg" : "bright-bg";
              return (
                <li className="plan-li" key={curPlan.id}>
                  <div className="plan-container">
                    <div className={`plan-header ${backgroundColor}`}>
                      <span className="plan-title">{curPlan.plan_title}</span>
                      <span className="plan-title-details">
                        {curPlan.plan_details}
                      </span>
                    </div>
                    {curPlan.plan_highlights.length > 0 ? (
                      <div className={`plan-highlight ${backgroundColor}`}>
                        <img
                          className="highlights-icon"
                          src={curPlan.plan_highlights[0].h_icon}
                        ></img>
                        <div style={{ display: "inline" }}>
                          <span className="bold">
                            {curPlan.plan_highlights[0].h_title}
                          </span>
                          <span>
                            {curPlan.plan_highlights[0].h_description}
                          </span>
                        </div>
                      </div>
                    ) : (
                      <Fragment />
                    )}
                    <div className="plan-advantages">
                      {curPlan.plan_advantages.map(advantage => {
                        return (
                          <div className="advantage">
                            <p>
                              <span className="bold">{advantage.a_title}</span>
                              <span>{advantage.a_description}</span>
                            </p>
                          </div>
                        );
                      })}
                    </div>
                    <div className="plan-end-section">
                      <Price
                        price={curPlan.plan_price}
                        days={curPlan.plan_limit_days}
                      />

                      <span
                        class="plan-small-words"
                        style={{ display: "block" }}
                      >
                        {`Válido por ${curPlan.plan_limit_days || 0} dias`} dias
                      </span>
                      <p class="plan-small-words">{curPlan.plan_small_words}</p>
                    </div>
                    <DetailsModal
                      button={
                        <div className="mt">
                          <SubscribeBtn>Assinar</SubscribeBtn>
                        </div>
                      }
                    ></DetailsModal>
                    <div className="plan-tray">
                      <DetailsModal
                        title="Detalhes do plano"
                        button={
                          <SubscribeBtn color="purple">
                            Ver Detalhes
                          </SubscribeBtn>
                        }
                      >
                        <div
                          dangerouslySetInnerHTML={{
                            __html: curPlan.plan_more_details.franchise
                          }}
                        ></div>
                      </DetailsModal>
                    </div>
                  </div>
                </li>
              );
            })}
          </ul>
        ) : (
          <Fragment />
        )}
      </Fragment>
    );
  }
}
