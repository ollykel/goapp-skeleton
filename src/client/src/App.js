import React, { Component } from 'react';
import '../public/App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src="/public/gopher.png" className="App-logo" alt="logo" />
          <p>Welcome to your first Goapp!</p>
		  <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
