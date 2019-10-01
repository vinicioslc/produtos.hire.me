import React, { Fragment, Component } from "react";
import "./DetailsModal.css";
export default class DetailsModal extends React.Component {
  state = { show: false };

  showModal = () => {
    if (this.props.children) {
      this.setState({ show: true });
      setTimeout(() => {
        this.setState({ openAnimation: true });
      }, 100);
    }
  };

  hideModal = () => {
    this.setState({ show: false });
    setTimeout(() => {
      this.setState({ openAnimation: false });
    }, 100);
  };

  render() {
    return (
      <main>
        <Modal
          show={this.state.show}
          animate={this.state.openAnimation}
          handleClose={this.hideModal}
          title={this.props.title}
          theme={this.props.theme}
        >
          {this.props.children}
        </Modal>
        <div className="modal-button" type="button" onClick={this.showModal}>
          {this.props.button}
        </div>
      </main>
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
        {title ? <div>{title}</div> : <Fragment></Fragment>}
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
