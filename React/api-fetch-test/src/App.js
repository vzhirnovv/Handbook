import React, { Component } from "react";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      posts: []
    };
  }

  componentDidMount() {
    fetch("https://jsonplaceholder.typicode.com/posts")
      .then(response => response.json())
      .then(result => this.setState({ posts: result }));
  }

  render() {
    return (
      <div>        
        <ul>
          {this.state.posts.map(function(item, key) {
            return (
              <li key={key}>
                <h1>{item.id}</h1>
                <h3>{item.title}</h3>
                <p>{item.body}</p>
              </li>
            );
          })}
        </ul>
      </div>
    );
  }
}

export default App;
