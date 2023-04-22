package ws

var BroadcastChannel = Channel{
	register:    make(chan *Client),
	unregister:  make(chan *Client),
	clientsPool: make(map[*Client]bool),
	send:        make(chan Message, 100)}
