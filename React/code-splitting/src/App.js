import React, { Component, Suspense } from 'react';
// import LazyComponent from './lazy.jsx';
import './App.css';

const LazyComponent = React.lazy(() => import('./lazy.jsx'))


class App extends Component {
   
  constructor(props) {
    super(props)
    this.state = {
      showLazyComponent: false
    };
  }

  togglePostList = () => {
    this.setState({
      showLazyComponent: !this.state.showLazyComponent
    })
  }

  render() { 
    {/* dynamic import example */}
    import("./math.js").then(Math => {
      console.log(Math.add(15, 25))
    });

    return (
      <div className="App">
	<button onClick={this.togglePostList}>Show post list!</button>
	{this.state.showLazyComponent && 
	    (
	      <Suspense fallback={<div>Loading...</div>}>
		<LazyComponent />       
	      </Suspense>
	    )
	}	
      </div>
    );
  }
}

export default App;
