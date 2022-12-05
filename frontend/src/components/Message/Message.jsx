import React, { useState } from "react";
import './Message.css'

const Message = ({ message }) => {
  let temp = JSON.parse(message)
  const [stateMessage, setStateMessage] = useState(temp)

  return (
    <div className="Message">
      {stateMessage.body}
    </div>
  )
}

export default Message
