import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

function App() {
  return (
    <Router>
      {/* Menu */}
      <Switch>
        <Route path="/" children={Home} />
        <Route path="/config" children={Config} />
      </Switch>
    </Router>
  );
}

// Dummys
const Home = () => "Home";
const Config = () => "Config";

export default App;
