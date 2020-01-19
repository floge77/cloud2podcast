import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import { ThemeProvider } from "@chakra-ui/core";
import customTheme from "./theme";

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

const withProvider = () => (
  <ThemeProvider theme={customTheme}>
    <App />
  </ThemeProvider>
);

export default withProvider;
