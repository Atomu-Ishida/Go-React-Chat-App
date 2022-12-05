package main

import (
	"fmt"
	"net/http"
	"github.com/Atomu-Ishida/Go-React-Chat-App/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request){
	fmt.Println("websocketの通信が完了しました")

	conn, err := websocket.Upgrade(w,r) // http通信をwebsocketにアップグレードする

	if err != nil{ // エラーハンドリング
		fmt.Fprintf(w, "%+V\n", err)
	}
	client := &websocket.Client{ // clientを作成
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client // pool.Resisterチャネルにclientを格納
	client.Read() //メッセージを読み取るメソッドを実行させる
}

func setupRoutes(){
	pool := websocket.NewPool() // ユーザーの状態を管理するPoolを初期化
	go pool.Start() // pool.Startメソッドをgoルーチンとして実行させる

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWS(pool, w, r) // サーバーを立てる
	})
}

func main()  {
	fmt.Println("リアルタイムチャットアプリを起動")
	setupRoutes()
	http.ListenAndServe(":9000",nil)
}