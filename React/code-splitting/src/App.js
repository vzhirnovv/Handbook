import React, { Suspense } from 'react';
import './App.css';

const LazyComponent = React.lazy(() => import('./lazy.jsx'))

function App() {
  
  {/* dynamic import example */}
  import("./math.js").then(Math => {
    console.log(Math.add(15, 25))
  });

  return (
    <div className="App">
      {/* Suspense example */}
      <Suspense fallback={<div>Loading...</div>}>
	<LazyComponent />       
      </Suspense>	
    </div>
  );
}

export default App;
