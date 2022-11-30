import React from "react";
import Header from './components/Header/Header'
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import './App.css'
import { connect, sendMsg } from './api'

const App = () => {
  const [chatHistory, setChatHistory] = React.useState([])

  React.useEffect(() => {
    connect((msg) => {
      console.log("New Message")
      setChatHistory((prevState) => ([...prevState, msg]))
      console.log(chatHistory)
    })
  }, [])

  const send = (event) => {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <ChatInput send={send} />
    </div>
  )
}

export default App
