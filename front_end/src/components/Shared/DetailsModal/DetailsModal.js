import React, { Fragment } from "react";
import plansApi from "../../../services/plans-api";
import ThemmedComponent from "../../Base/ThemmedComponent";
import "./DetailsModal.css";
export default class DetailsModal extends ThemmedComponent {
  state = { show: false };

  fetchDetails = () => {
    this.setState({
      details: undefined
    });
    console.log("downloaded data");
    plansApi
      .getPlanDetails({ carrier: this.getTheme(), plan_sku: this.props.sku })
      .then(result => {
        this.setState({
          details: result.franchise
        });
        console.log("downloaded data", result);
      })
      .catch(reas => {
        console.error(reas);
        this.setState({
          details: reas.toString()
        });
      });
  };

  showModal = () => {
    this.fetchDetails();
    this.setState({ show: true });
    setTimeout(() => {
      this.setState({ openAnimation: true });
    }, 100);
  };

  hideModal = () => {
    this.setState({ show: false });
    setTimeout(() => {
      this.setState({ openAnimation: false });
    }, 100);
  };

  render() {
    return (
      <div>
        <Modal
          show={this.state.show}
          animate={this.state.openAnimation}
          handleClose={this.hideModal}
          title={this.props.title}
          theme={this.props.theme}
        >
          <div
            dangerouslySetInnerHTML={{
              __html: this.state.details
            }}
          ></div>
        </Modal>
        <div className="modal-button" type="button" onClick={this.showModal}>
          {this.props.button}
        </div>
      </div>
    );
  }
}

const Modal = ({ title, handleClose, show, children, animate, theme }) => {
  const showHideClassName = show ? "modal display-block" : "modal display-none";

  return (
    <div onClick={handleClose} className={showHideClassName}>
      <section
        onClick={e => e.stopPropagation()}
        className={`modal-main ${
          animate ? "animation-open" : "animation-closed"
        } ${theme}`}
      >
        {title ? (
          <div>
            <div className={`modal-title ${theme}`}>{title}</div>
          </div>
        ) : (
          <Fragment></Fragment>
        )}
        <button
          className={`close-btn primary-text-color ${theme}`}
          onClick={handleClose}
        >
          X
        </button>
        {children}
      </section>
    </div>
  );
};
