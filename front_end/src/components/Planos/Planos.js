import React, { Fragment, Component } from "react";
import "./Planos.css";
export default class Planos extends Component {
  constructor() {
    super();
    this.state = {
      plans: [
        {
          id: "5d8fb4a002eb811ad881f601",
          plan_carrier: "vivo",
          plan_title: "3GB",
          plan_details: " de internet 4.5G",
          plan_small_words: "com renovação automática",
          plan_price: 15.99,
          plan_highlights: [
            {
              h_icon: "whattsicon.png",
              h_title: "WhatsApp Ilimitado",
              h_description: " para mensagens, vídeos e fotos"
            }
          ],
          plan_advantages: [
            {
              a_title: "Ganhe 1GB de bônus, ",
              a_description: "válido por 7 dias, ao acumular R$35 em recarga"
            },
            {
              a_title: "Minutos ilimitados ",
              a_description:
                "em ligações locais para celulares de outras operadoras"
            },
            {
              a_title: "Ligações ilimitadas ",
              a_description:
                "para celulares e fixos Vivo de todo o Brasil, usando o 15"
            },
            {
              a_title: "SMS ilimitado ",
              a_description: "para celulares Vivo"
            },
            {
              a_title: "Serviços Digitais: ",
              a_description: "NBA, GoRead e Vivo Bem"
            }
          ],
          plan_limit_mbytes: 2000,
          plan_limit_days: 7
        },
        {
          id: "5d8fb4a002eb8191205192d881f601",
          plan_carrier: "vivo",
          plan_title: "1.5GB",
          plan_details: "de internet 4.5G",
          plan_small_words: "com renovação automática",
          plan_price: 15.99,
          plan_highlights: [
            {
              h_icon: "whattsicon.png",
              h_title: "WhatsApp Ilimitado",
              h_description: "para mensagens, vídeos e fotos"
            }
          ],
          plan_advantages: [
            {
              a_title: "Ganhe 1GB de bônus, ",
              a_description: "válido por 7 dias, ao acumular R$35 em recarga"
            },
            {
              a_title: "Minutos ilimitados ",
              a_description:
                "em ligações locais para celulares de outras operadoras"
            }
          ],
          plan_limit_mbytes: 2000,
          plan_limit_days: 7
        },
        {
          id: "5d8fb4a002eb8191205192d881f601",
          plan_carrier: "vivo",
          plan_title: "200MIN",
          plan_details: "",
          plan_small_words: "com renovação automática",
          plan_price: 15.99,
          plan_highlights: [],
          plan_advantages: [
            {
              a_title: "Ganhe 1GB de bônus, ",
              a_description: "válido por 7 dias, ao acumular R$35 em recarga"
            },
            {
              a_title: "Minutos ilimitados ",
              a_description:
                "em ligações locais para celulares de outras operadoras"
            }
          ],
          plan_limit_mbytes: 2000,
          plan_limit_days: 7
        }
      ]
    };
  }

  render() {
    return (
      <Fragment>
        <h1 className="offer-title">Conheça as Ofertas</h1>
        <ul>
          {this.state.plans.map(curPlan => {
            return (
              <li className="card-li">
                <div className="card-container">
                  <div className="card-header">
                    <span className="card-title">{curPlan.plan_title}</span>
                    <small className="card-title-details">
                      {curPlan.plan_details}
                    </small>
                  </div>
                  {curPlan.plan_highlights.length > 0 ? (
                    <div className="card-highlight">
                      <span className="bold">
                        {curPlan.plan_highlights[0].h_title}
                      </span>
                      <span>{curPlan.plan_highlights[0].h_description}</span>
                    </div>
                  ) : (
                    <Fragment></Fragment>
                  )}
                  <div className="card-advantages">
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
                </div>
              </li>
            );
          })}
        </ul>
      </Fragment>
    );
  }
}
