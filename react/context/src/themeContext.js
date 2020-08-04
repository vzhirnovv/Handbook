// Global app context is defined here with a default value of 'none'.
// After that, the ThemeProvider and Consumer are exported into the project.

import React from "react";

const ThemeContext = React.createContext("none");

export const ThemeProvider = ThemeContext.Provider;

export const ThemeConsumer = ThemeContext.Consumer;
