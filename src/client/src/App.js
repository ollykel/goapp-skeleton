import React, { Component } from 'react';
import '../public/App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src="gopher.png?v=2" className="App-logo" alt="logo" />
          <p>Welcome to your first Goapp!</p>
		  <a
            className="App-link"
            href="https://github.com/ollykel/webapp"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn about Goapp
          </a>
        </header>
      </div>
    );
  }
}

export default App;
