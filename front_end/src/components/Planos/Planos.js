import React, { Fragment, Component } from "react";
import Price from "./Price/Price";
import SubscribeBtn from "./SubscribeBtn/SubscribeBtn";
import DetailsModal from "../DetailsModal/DetailsModal";
import "./Planos.css";
export default class Planos extends Component {
  constructor(props) {
    super(props);
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
              h_icon:
                "https://celular.vivo.com.br/planos/pre/img/whattsicon.png",
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
          plan_price: 15.89,
          plan_highlights: [
            {
              h_icon:
                "https://celular.vivo.com.br/planos/pre/img/whattsicon.png",
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
            }
          ],
          plan_limit_mbytes: 2000,
          plan_limit_days: 7
        },
        {
          id: "5d8fb4a002eb846205192d881f601",
          plan_carrier: "vivo",
          plan_title: "200MIN",
          plan_details: "",
          plan_small_words: "com renovação automática",
          plan_price: 7.99,
          plan_highlights: [
            {
              h_title: "Ligações Ilimitadas",
              h_description: " para outro vivo."
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
        }
      ].map(val => {
        val.plan_more_details = {
          franchise: "",
          additional_info: `A velocidade de referência de download é de até 5 Mbps e de upload de até 500Kbps.

          Caso a franquia de internet seja 100% consumida antes do término dos 7 dias de validade do pacote, o serviço será interrompido. Para voltar a navegar o cliente poderá antecipar renovação da sua promoção Vivo Pré. Para antecipar a renovação, basta ter saldo de R$14,99 e enviar um SMS com a palavra ANTECIPE para 9003. A promoção será renovada por mais 7 dias e o cliente receberá todos os benefícios novamente.
          
          Para a navegação na internet, com velocidade 4G, de até 5Mbps para download e até 500Kbps para upload, é necessária a utilização de um chip e um smartphone LTE (frequência 2.5 GHz) compatíveis com a tecnologia 4G. Nas localidades que não possuírem cobertura 4G a alteração de rede para 3G é feita automaticamente. Para consultar a cobertura 3G e 4G acesse www.vivo.com.br/cobertura. Caso o cliente não possua aparelho ou chip compatível com 4G, a velocidade de navegação na internet é de até 1Mbps para download e 25Kbps para upload.
          
          Ao atingir 100% da franquia de minutos para outras operadoras o cliente será tarifado conforme o seu plano. Para mais informações clique aqui.
    
          Ligações locais para telefones fixos de outras operadoras por R$0,69 (sessenta e nove centavos) por minuto iniciado.`,
          regulation: ""
        };
        return val;
      })
    };
  }

  render() {
    return (
      <Fragment>
        <h1 className="offer-title">Ofertas feitas para você</h1>
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
                        <span>{curPlan.plan_highlights[0].h_description}</span>
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
                    <Price value={curPlan.plan_price} />
                    <SubscribeBtn></SubscribeBtn>
                    <p class="plan-small-words">{curPlan.plan_small_words}</p>
                  </div>

                  <div className="plan-tray">
                    <DetailsModal buttonName="Mais Informações">
                      <div>{curPlan.plan_more_details.additional_info}</div>
                      <div>{curPlan.plan_more_details.additional_info}</div>
                    </DetailsModal>
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
