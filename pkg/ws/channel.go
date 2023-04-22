package ws

type Channel struct {
	register    chan *Client
	unregister  chan *Client
	clientsPool map[*Client]bool
	send        chan Message
}

func (channel *Channel) ProcessTask() {
	for {
		select {
		case client := <-channel.register:
			channel.clientsPool[client] = true

		case client := <-channel.unregister:
			if channel.clientsPool[client] {
				delete(channel.clientsPool, client)
			}

		case message := <-channel.send:
			for client := range channel.clientsPool {
				select {
				case client.send <- message:
					// nothing to do here
				default:
					break
				}
			}
		}
	}
}

func (channel *Channel) Broadcast(message Message) {
	// If chan is full, skip
	select {
	case BroadcastChannel.send <- message:
	default:
		break
	}
}
