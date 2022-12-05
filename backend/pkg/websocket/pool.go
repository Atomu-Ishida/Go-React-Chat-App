package websocket

import(
	"fmt"
)

type Pool struct { // ユーザーの入室/退室/チャット送信を管理する構造体
	Register chan *Client // ユーザーの入室時に利用するチャネル
	Unregister chan *Client // ユーザーの退室時に利用するチャネル
	Clients map[*Client]bool // チャットの参加者一覧
	Broadcast chan Message // チャットをやりとりするためのチャネル
}

func NewPool() *Pool{ // 初期化用の関数
	return &Pool{
		Register: make(chan *Client),
		Unregister: make(chan *Client),
		Clients: make(map[*Client]bool),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start(){
	for{ // for文で無限ループを作る
		select{ // チャネルは値が入っていない場合、受信をブロックするがブロックせずに使いためselectを利用する
		case client := <-pool.Register: // clientがpool.Resisterを受信した場合実行
			pool.Clients[client] = true // Clientsに追加する
			fmt.Println("現在、入室している人数は:", len(pool.Clients), "人です") // pool.Clientsのlenで入室している人数を表示する
			for client, _ := range pool.Clients{
				client.Conn.WriteJSON(Message{Type: 1, Body: "新しいユーザーが入室しました"}) // 現在入室しているユーザー全てにメッセージを送る
			}
			break
		case client := <-pool.Unregister: // clientがpool.Unregisterを受信した場合に実行
			delete(pool.Clients, client) // deleteメソッドでpool.Clientsマップからclientを削除
			fmt.Println("現在、入室している人数は:", len(pool.Clients), "人です") // pool.Clientsのlenで入室している人数を表示する
			for client, _ := range pool.Clients{
				client.Conn.WriteJSON(Message{Type: 1, Body: "ユーザーが退出しました"}) // 現在入室しているユーザー全てにメッセージを送る
			}
			break
		case message := <-pool.Broadcast: // messageがpool.Broadcastを受信した場合に実行
			fmt.Println("入室しているユーザーにメッセージを送りました")
			for client, _ := range pool.Clients{
				if err := client.Conn.WriteJSON(message); err != nil{ // 受信したmessageをすべてのユーザーに送る
					fmt.Println(err)
					return
				}
			}
		}
	}
}