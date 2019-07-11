import React, { Component } from "react";
import ErrorBoundary from "./errorBoundary";

const CheckList = props => {
    if (props.name === "error") {
        throw new Error("Error Here!");
    }
    return <h1>{props.name}</h1>;
};

class App extends Component {
    render() {
        return (
            <div>
                <ErrorBoundary>
                    <CheckList name="ok" />
                </ErrorBoundary>
                <ErrorBoundary>
                    <CheckList name="error" />
                </ErrorBoundary>
                <ErrorBoundary>
                    <CheckList name="very ok" />
                </ErrorBoundary>
                <ErrorBoundary>
                    <CheckList name="really nice" />
                </ErrorBoundary>
                <h1>Error boundary test.</h1>
            </div>
        );
    }
}

export default App;
