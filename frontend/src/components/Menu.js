import React from "react";
import styled from "@emotion/styled";
import { Link } from "react-router-dom";
import { Box } from "@chakra-ui/core";

const _Menu = (props) => (
  <Box className={props.className} bg="brand.900" color="white">
    <header>
      <nav>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/config">Config</Link>
          </li>
        </ul>
      </nav>
    </header>
  </Box>
);

const Menu = styled(_Menu)`
  ul {
    display: flex;
    list-style: none;
    margin: 0 auto;
    padding: 10px 0;
  }

  a {
    color: white;
    text-decoration: none;
    padding: 10px;
  }
`;

export default Menu;
