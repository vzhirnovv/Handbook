import React, { Component } from "react";
import { BrowserRouter, Route, Link } from "react-router-dom";
import { Hello } from "./Hello";
import "./App.css";

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Route exact path="/hello" component={Hello} />
          <Route path="/bye" component={Bye} />
          <Route path="/gayshit" component={Gayshit} />

          <hr />

          <ul>
            <li>
              <Link to="/hello">Welcome</Link>
            </li>
            <li>
              <Link to="/bye">Goodbye</Link>
            </li>
            <li>
              <Link to="/gayshit">Some gay fucking shit man</Link>
            </li>
          </ul>

          <hr />
        </div>
      </BrowserRouter>
    );
  }
}

const Bye = () => (
  <div>
    <h1>Aww why are you leaving don't :(((</h1>
  </div>
);

const Gayshit = () => (
  <div>
    <h1>
      Btw have i told you im on some gay shit its called supreme man its so c o
      o l
    </h1>
  </div>
);

export default App;
