package ws

import (
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketHandler struct {
	wsUpgrader websocket.Upgrader
}

func NewWebSocketHandler() (*WebSocketHandler, error) {
	return &WebSocketHandler{
		wsUpgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  4096, // 4k
			WriteBufferSize: 4096, // 4k
		}}, nil
}

func (wsh *WebSocketHandler) HandleConn(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Errorf("[WS] Upgrade failed: %s", err.Error())
		return
	}

	client := Client{
		conn: conn,
		send: make(chan Message, 100),
	}

	go client.RecvMessage()
	go client.SendMessage()

	BroadcastChannel.register <- &client
}
