import React from "react";

import Planos from "./components/Planos/Planos";
import logo from "./logo.svg";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <a className="LogoHeader"></a>
      </header>
      <div>
        <Planos> </Planos>
      </div>
    </div>
  );
}

export default App;
