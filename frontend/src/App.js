import React, { Component } from "react";
import "./App.css";
import { connect, send_msg } from "./api";
import Header from './components/Header/Header.jsx';
import ChatHistory from './components/Chathistory/ChatHistory.jsx'
import Input from './components/Input/Input.jsx'



class App extends Component {
	constructor(props) {
	super(props);
	this.state = {
		chatHistory: []
		}
	}
	componentDidMount() {
		connect((msg) => {
		console.log("New Message")
		this.setState(prevState => ({
			chatHistory: [...this.state.chatHistory, msg]
			}))
			console.log(this.state);
		});
	}
	send(event) {
		if(event.keyCode===13){
			send_msg(event.target.value);
			event.target.value = "";
		}
	}

  render() {
    return (
      <div className="App">
		<Header />
		<ChatHistory chatHistory={this.state.chatHistory} />
		<Input send={this.send} />
      </div>
    );
  }
}

export default App;

