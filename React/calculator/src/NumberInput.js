import React, {Component} from 'react';

const numberOrder = {
    1: 'first',
    2: 'second'
};

export class NumberInput extends Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(e) {
        this.props.handleValueChange(e.target.value);
    }

    render() {
        const number = this.props.value;
        const order = this.props.order;
        return (
            <fieldset>
                <legend>Enter the {numberOrder[order]} number</legend>
                <input value={number} onChange={this.handleChange}/>
            </fieldset>
        );
    }
}
