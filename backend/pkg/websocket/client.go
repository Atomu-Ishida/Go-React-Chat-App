package websocket

import(
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn // websocketのconnectionの参照
	Pool *Pool // Poolの参照
}

type Message struct { // クライアントに送るメッセージの構造体
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read(){ // メッセージを読み取るメソッド
	defer func(){ // 接続が切断されると実行される
		c.Pool.Unregister <- c // c.Pool.Unregisterチャネルにcを格納
		c.Conn.Close() // *Clientの持つwebsocketコネクションをCloseする
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage() // ReadMessageでメッセージを受信する
		if err != nil{
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p)} // messageにMessage構造体を格納・buffer型のpをstringに型変換する
		c.Pool.Broadcast <- message // c.Pool.Broadcastチャネルにmessageを格納
		fmt.Printf("メッセージを受け取りました: %+v\n", message)
	}
}