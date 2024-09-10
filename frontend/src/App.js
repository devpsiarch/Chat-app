import React, { Component } from "react";
import "./App.css";
import { connect, send_msg } from "./api";
import Header from './components/Header/Header.jsx';

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    send_msg("hello");
  }

  render() {
    return (
      <div className="App">
		<Header />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;

