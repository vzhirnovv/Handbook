import React from "react";
import HigherOrderComponent from "./HOC";

const SomeComponent = props => {
    return <h1>some component</h1>;
};

const WrappedComponent = HigherOrderComponent(SomeComponent);

function App() {
    return (
        <div className="App">
            <SomeComponent />
            <WrappedComponent />
        </div>
    );
}

export default App;
