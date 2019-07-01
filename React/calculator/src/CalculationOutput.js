import React, {Component} from 'react';

export class CalculationOutput extends Component {

    add(v1, v2) {
        return v1 + v2;
    }

    subtract(v1, v2) {
        return v1 - v2;
    }

    multiply(v1, v2) {
        return v1 * v2;
    }

    divide(v1, v2) {
        return v1 / v2;
    }

    render() {
        return (
            <div>
                <p>The result of adding these two numbers is {this.add(this.props.value1, this.props.value2)}</p>
                <p>The result of subtracting these two numbers is {this.subtract(this.props.value1, this.props.value2)}</p>
                <p>The result of multiplying these two numbers is {this.multiply(this.props.value1, this.props.value2)}</p>
                <p>The result of dividing these two numbers is {this.divide(this.props.value1, this.props.value2)}</p>
            </div>
        );
    }
}


