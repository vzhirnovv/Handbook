import React, {Component} from 'react';
import './App.css';

class Calculator extends Component {
    constructor(props) {
        super(props);
        this.value1 = this.props.value1;
        this.value2 = this.props.value2;
    }

    add() {
        return this.props.value1 + this.props.value2
    }

    subtract() {
        return this.props.value1 - this.props.value2
    }

    multiply() {
        return this.props.value1 * this.props.value2
    }

    divide() {
        return this.props.value1 / this.props.value2
    }

    render() {
        return (
            <div>
                Adding {this.props.value1} and {this.props.value2} will result in {this.add()}
                <br/>
                <br/>
                Subtracting {this.props.value1} from {this.props.value2} will result in {this.subtract()}
                <br/>
                <br/>
                Multiplying {this.props.value1} and {this.props.value2} will result in {this.multiply()}
                <br/>
                <br/>
                Dividing {this.props.value1} by {this.props.value2} will result in {this.divide()}
            </div>
        );
    }
}

export default Calculator;
