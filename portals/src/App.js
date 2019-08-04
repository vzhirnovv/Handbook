import React, { Component } from "react";
import "./App.css";

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            showModal: false
        };
    }

    toggleModal = () => {
        this.setState({
            showModal: !this.state.showModal
        });
    };
}

export default App;
