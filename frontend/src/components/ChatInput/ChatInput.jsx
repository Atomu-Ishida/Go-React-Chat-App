import React from "react";
import './ChatInput.css'


const ChatInput = ({ send }) => { // propsでチャットを送るための関数を受け取る
  return (
    <div className="ChatInput">
      <input onKeyDown={send} placeholder="何かメッセージを入力してください 入力したらエンターを押してください"></input>
    </div>
  )
}

export default ChatInput