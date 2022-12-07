import React from "react";
import './ChatHistory.css'
import Message from '../Message/Message'

const ChatHistory = ({ chatHistory }) => { // propsでチャット履歴を受け取る
  console.log(chatHistory)
  const messages = chatHistory.map(msg => <Message key={msg.timeStamp} message={msg.data} />)

  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {messages}
    </div>
  )
}

export default ChatHistory