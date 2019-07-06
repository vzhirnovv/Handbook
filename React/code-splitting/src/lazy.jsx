import React, { Component } from 'react'

class LazyComponent extends Component { 
  fetchData = async () => {
    await new Promise(resolve => setTimeout(resolve, 3000))
    const response = await fetch('https://jsonplaceholder.typicode.com/posts')
    const result = await response.json()
    console.log(result)
    this.setState({
      data: result
    })
    console.log('data acquired!')
}
  constructor(props) {
    super(props)

    this.state = {
      data: [] 
    }
  }

  componentDidMount() {
    this.fetchData() 
    console.log(this.state.data) 
  }

  render() {
    return (
      <div>
	<h1>Lazy component loaded!</h1>
	<ul>
	  {this.state.data.map((item, key) => {
	    return (
	      <li key={key}>
		<h3>{item.userId}</h3>
		<h4>{item.title}</h4>
		<p>{item.body}</p>
	      </li>
	    )
	  })}
	</ul>
      </div>
    )
  }
}

export default LazyComponent
