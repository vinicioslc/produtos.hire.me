import React, { Fragment, Component } from "react";
import "./DetailsModal.css";
export default class ModalButton extends React.Component {
  state = { show: false };

  showModal = () => {
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
      <main>
        <Modal
          show={this.state.show}
          animate={this.state.openAnimation}
          handleClose={this.hideModal}
        >
          {this.props.children}
        </Modal>
        <button className="modal-button" type="button" onClick={this.showModal}>
          {this.props.buttonName}
        </button>
      </main>
    );
  }
}

const Modal = ({ handleClose, show, children, animate }) => {
  const showHideClassName = show ? "modal display-block" : "modal display-none";

  return (
    <div className={showHideClassName}>
      <section
        className={`modal-main ${
          animate ? "animation-open" : "animation-closed"
        }`}
      >
        <button className="close-btn purple-text" onClick={handleClose}>
          X
        </button>
        {children}
      </section>
    </div>
  );
};
