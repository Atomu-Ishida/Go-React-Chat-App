var socket = new WebSocket('ws://localhost:9000/ws') // Websocketオブジェクトを作成

let connect = (cb) => {
  console.log("接続しています")

  socket.onopen = () => {
    console.log("接続完了！")
  }

  socket.onmessage = (msg) => { // websocketを通してメッセージを受け取った時に実行
    console.log("メッセージを受信しました", msg)
    cb(msg)
  }

  socket.onclose = (event) => {
    console.log("接続を切断します", event)
  }

  socket.onerror = (error) => {
    console.log("エラーが起きました:", error)
  }
}

let sendMsg = (msg) => {
  console.log("メッセージを送ります:", msg)
  socket.send(msg)
}

export { connect, sendMsg }