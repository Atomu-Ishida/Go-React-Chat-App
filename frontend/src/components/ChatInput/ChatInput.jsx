import React, { Component } from "react";
import './ChatInput.css'

class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput">
        <input onKeyDown={this.props.send} placeholder="何かメッセージを入力してください 入力したらエンターを押してください"></input>
      </div>
    )
  }
}

export default ChatInput