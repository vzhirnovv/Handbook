import React from "react";
import { ThemeConsumer, ThemeProvider } from "./themeContext";

function App() {
    return (
        <div className="App">
            {/* every child of <Toolbar> can now access ThemeContext */}
            {/* the new value is passed to the context when a provider is used */}
            {/* Every Context object comes with a Provider React component that allows consuming components to subscribe to context changes.
                Accepts a value prop to be passed to consuming components that are descendants of this Provider. One Provider can be connected to many consumers. 
                Providers can be nested to override values deeper within the tree.
                All consumers that are descendants of a Provider will re-render whenever the Provider’s value prop changes. 
                The propagation from Provider to its descendant consumers is not subject to the shouldComponentUpdate method, 
                so the consumer is updated even when an ancestor component bails out of the update. */}
            <ThemeProvider value="light">
                <Toolbar />

                {/* ThemeProvider is overridden - dark will be shown */}
                <ThemeProvider value="dark">
                    <Toolbar />
                </ThemeProvider>

                <Toolbar />
            </ThemeProvider>
            <ThemeProvider value="dark">
                <Toolbar />
            </ThemeProvider>
            <ThemeProvider value="dark">
                <Toolbar />
            </ThemeProvider>
        </div>
    );
}

const Toolbar = props => {
    return (
        <div>
            <ThemedButton />
        </div>
    );
};

const ThemedButton = props => {
    return (
        // ThemedButton is the child of Toolbar, which means it can access ThemeContext via ThemeConsumer
        // Theme consumer must have a function as a child. This function recieves the current context value
        // and returns a node.
        <ThemeConsumer>{context => <button>{context}</button>}</ThemeConsumer>
    );
};

export default App;
