var socket = new WebSocket('ws://localhost:9000/ws')

let connect = (cb) => {
  console.log("connecting")

  socket.onopen = () => {
    console.log("successfully connected")
  }

  socket.onmessage = (msg) => {
    console.log("Message from websocket:", msg)
  }

  socket.onclose = (event) => {
    console.log("socket close connection:", event)
  }

  socket.onerror = () => {
    console.log("socket error:", error)
  }
}

let sendMsg = (msg) => {
  console.log("sending msg:", meg)
  socket.send(msg)
}

export default { connect, sendMsg }