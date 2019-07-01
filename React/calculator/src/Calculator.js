import React, {Component} from 'react';
import {NumberInput} from "./NumberInput";
import {CalculationOutput} from "./CalculationOutput";




export class Calculator extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value1: '',
            value2: ''
        };

        this.handleFirstValue = this.handleFirstValue.bind(this);
        this.handleSecondValue = this.handleSecondValue.bind(this)
    }

    handleFirstValue(value1) {
        this.setState({
            value1
        })
    }

    handleSecondValue(value2) {
        this.setState({
            value2
        })
    }


    render() {
        const v1 = this.state.value1;
        const v2 = this.state.value2;

        return (
            <div>
                <NumberInput order={1}
                             handleValueChange={this.handleFirstValue}
                />
                <NumberInput order={2}
                             handleValueChange={this.handleSecondValue}
                />
                <p>the numbers you entered are {parseInt(v1, 10)} and {parseInt(v2, 10)}</p>
                <CalculationOutput value1={parseInt(v1, 10)} value2={parseInt(v2, 10)} />
            </div>
        );
    }
}


