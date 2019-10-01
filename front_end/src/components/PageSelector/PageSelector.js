import React, { Fragment, Component } from "react";
import "./PageSelector.css";
import { BrowserRouter as Router, Link } from "react-router-dom";

export default class PageSelector extends Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <Fragment>
        <div className="selector-container">
          <Link to="/vivo">
            <button className={`selector-btn vivo left-rounded`}>
              <span className="">Vivo</span>
            </button>
          </Link>
          <Link to="/claro">
            <button className={`selector-btn claro right-rounded `}>
              <span className="">Claro</span>
            </button>
          </Link>
        </div>
      </Fragment>
    );
  }
}
