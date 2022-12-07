import React, { useState, useEffect } from "react";
import Header from './components/Header/Header'
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import './App.css'
import { connect, sendMsg } from './api'

const App = () => {
  const [chatHistory, setChatHistory] = useState([]) // チャット履歴をstateで管理

  useEffect(() => { // useEffect内でstate更新をする
    connect((msg) => {
      console.log("新しいメッセージです")
      setChatHistory([...chatHistory, msg])
      console.log(chatHistory)
    })
  }, [chatHistory])

  const send = (event) => { // チャットを送る関数
    if (event.keyCode === 13) { // エンターキーが押されたときに実行
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
