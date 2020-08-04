// Error boundaries are React components that catch JavaScript errors anywhere in their child component tree,
// log those errors, and display a fallback UI instead of the component tree that crashed.
// Error boundaries catch errors during rendering, in lifecycle methods,
// and in constructors of the whole tree below them.

import React, { Component } from "react";

export default class ErrorBoundary extends Component {
    constructor(props) {
        super(props);
        this.state = {
            hasError: false
        };
    }

    // A class component becomes an error boundary if it defines either (or both)
    // of the lifecycle methods static getDerivedStateFromError() or componentDidCatch().
    // Use static getDerivedStateFromError() to render a fallback UI after an error has been thrown.
    // Use componentDidCatch() to log error information.

    static getDerivedStateFromError(error) {
        // Update state so the next render will show the fallback UI.
        return {
            hasError: true
        };
    }

    componentDidCatch(error, info) {
        // You can also log the error to an error reporting service
        console.log("caught error!");
        console.log(error.toString());
        console.log(info);
        this.setState({
            hasError: true
        });
    }

    render() {
        if (this.state.hasError) {
            return <h1>Error caught.</h1>;
        }
        return this.props.children;
    }
}
