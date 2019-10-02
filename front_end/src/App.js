import React from "react";

import { Switch, Route, BrowserRouter as Router } from "react-router-dom";

import PageSelector from "./components/PageSelector/PageSelector";

import Vivo from "./components/Vivo/Index/Index.js";
import Claro from "./components/Claro/Index/Index.js";
import MainPage from "./components/Base/MainPage/MainPage";
import "./App.css";

function App() {
  return (
    <Router>
      <Switch>
        <Route
          path="/vivo"
          render={props => (
            <MainPage {...props} component={Vivo} title="Vivo" />
          )}
        />
        <Route
          path="/claro"
          render={props => (
            <MainPage {...props} component={Claro} title="Claro" />
          )}
        />
        <Route path="/" component={PageSelector} />
      </Switch>
    </Router>
  );
}

export default App;
