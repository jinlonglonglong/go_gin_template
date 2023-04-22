package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan Message
}

func (c *Client) RecvMessage() {
	//previous := int64(0)
	for {
		message := &Message{}

		if err := c.conn.ReadJSON(message); err != nil {
			c.conn.Close()
			BroadcastChannel.unregister <- c
			return
		}

		// limit sending message rate
		//now := time.Now().Unix()
		//if now - previous < 1 {
		//	continue
		//}
		//previous = now
		//BroadcastChannel.send <- *message
	}
}

func (c *Client) SendMessage() {
	for {
		message := <-c.send
		if err := c.conn.WriteJSON(message); err != nil {
			c.conn.Close()
			BroadcastChannel.unregister <- c
			return
		}
	}
}
