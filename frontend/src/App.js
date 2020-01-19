import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { ThemeProvider } from "@chakra-ui/core";
import customTheme from "./theme";
import Menu from "./components/Menu";

console.log(customTheme);

// Dummys
const Home = () => "Home";
const Config = () => "Config";

function App() {
  return (
    <Router>
      <Menu />

      <Switch>
        <Route exact path="/" children={Home} />
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
