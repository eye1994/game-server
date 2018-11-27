package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func onCreateEvent(conn *websocket.Conn, msg SocketEvent) {
	room, err := createRoom()
	if err != nil {
		log.Printf("Failed to create room: %v\n", err)
		return
	}

	client := &Client{conn: conn, room: room}
	client.send(SocketEvent{EventType: CreateEvent, Room: room.name})
	room.addClient(client, msg.PlayerName)
	handleGameData(client)
}

func onJoinEvent(conn *websocket.Conn, msg SocketEvent) {
	if rooms[msg.Room] == nil {
		log.Printf("Attempt to join an invalid room -> %v\n", msg)
		return
	}

	client := &Client{conn: conn, room: rooms[msg.Room]}
	rooms[msg.Room].addClient(client, msg.PlayerName)
	handleGameData(client)
}

func handleGameData(client *Client) {
	for {
		var msg SocketEvent
		err := client.conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		switch msg.EventType {
		case DataEvent:
			client.room.brodcast(client, msg)
		}
	}
}

// HandleWsConnection with()
func HandleWsConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

loop:
	for {
		var msg SocketEvent
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		switch msg.EventType {
		case CreateEvent:
			onCreateEvent(conn, msg)
			break loop
		case JoinEvent:
			onJoinEvent(conn, msg)
			break loop
		}
	}
}
