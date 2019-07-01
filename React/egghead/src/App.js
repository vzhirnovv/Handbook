import React from 'react';

class App extends React.Component {
    render() {
        return (
            <div>
                <Button>React</Button>
            </div>
        );
    }
}


const Button = (props) =>
    <button>{props.children}</button>

export default App