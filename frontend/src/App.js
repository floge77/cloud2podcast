import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { ThemeProvider } from "@chakra-ui/core";
import customTheme from "./theme";

import Menu from "./components/Menu";
import Home from "./components/Home";

console.log(customTheme);

// Dummys
const Config = () => "Config";

function App() {
  return (
    <Router>
      <Menu />

      <Switch>
        <Route exact path="/">
          <Home />
        </Route>
        <Route path="/config" children={Config} />
      </Switch>
    </Router>
  );
}

const withProvider = () => (
  <ThemeProvider theme={customTheme}>
    <App />
  </ThemeProvider>
);

export default withProvider;
