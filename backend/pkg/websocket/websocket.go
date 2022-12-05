package websocket

import(
	"net/http"
	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{ // 書き込みと読み込みのバッファサイズを指定する
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error){ // http通信をwebsocketにアップグレードするための関数
	upgrader.CheckOrigin = func(r *http.Request) bool {return true} // CheckOriginでCross-OriginをTrueにする
	conn, err := upgrader.Upgrade(w, r, nil) // Upgradeメソッドでhttp通信をwebsocketにアップグレードする
	if err != nil { // エラーハンドリング
		log.Println(err)
		return nil, err
	}
	return conn, nil
}