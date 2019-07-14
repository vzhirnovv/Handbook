import React, { Component } from "react";

const HigherOrderComponent = WrappedComponent => {
    class HOC extends Component {
        componentDidMount() {
            console.log("A wrapped component has just rendered!");
        }

        render() {
            return <WrappedComponent />;
        }
    }

    return HOC;
};

export default HigherOrderComponent;
