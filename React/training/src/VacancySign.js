import React, {Component} from "react";

export class VacancySign extends Component {
    render() {
        return (
            <div>
                <h1>This room is {this.props.vacancy ? 'vacant' : 'not vacant'} </h1>
            </div>
        );
    }
}
