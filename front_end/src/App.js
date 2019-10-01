import React from "react";

import { Switch, Route, BrowserRouter as Router } from "react-router-dom";

import PageSelector from "./components/PageSelector/PageSelector";

import Vivo from "./components/Vivo/Index/Index.js";
import Claro from "./components/Claro/Index/Index.js";
import "./App.css";

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/vivo" component={Vivo} />
        <Route path="/claro" component={Claro} />
        <Route path="/" component={PageSelector} />
      </Switch>
    </Router>
  );
}

export default App;
