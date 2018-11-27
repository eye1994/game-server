package ws

import (
	"github.com/gorilla/websocket"
)

const (
	// CreateEvent is the event name for creating a new game room
	CreateEvent = "create"
	// JoinEvent is the event for joining an existing game room
	JoinEvent = "join"
	// DataEvent is the event for sending data inside a game room
	DataEvent = "data"
	// JoinedEvent is the event for emiting a join event type
	JoinedEvent = "joined"
)

// RegisterUser with()
type RegisterUser struct {
	client     *Client
	playerName string
}

// SocketEvent with ()
type SocketEvent struct {
	EventType  string `json:"eventType"`
	PlayerName string `json:"playerName"`
	Room       string `json:"room"`
	Data       string `json:"data"`
}

// Client with()
type Client struct {
	conn *websocket.Conn
	room *Room
}

func (c *Client) send(data SocketEvent) {
	c.conn.WriteJSON(data)
}
