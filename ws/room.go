package ws

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

// Room with()
type Room struct {
	name    string
	clients map[*Client]string
}

func (r *Room) addClient(c *Client, playerName string) {
	r.clients[c] = playerName
	r.brodcast(c, SocketEvent{EventType: JoinedEvent, PlayerName: playerName})
}

func (r *Room) brodcast(from *Client, data SocketEvent) {
	for client := range r.clients {
		if client != from {
			client.send(data)
		}
	}
}

var rooms map[string]*Room

func init() {
	rooms = make(map[string]*Room)
	fmt.Printf("Created hub list -> %v\n", rooms)
}

func createRoom() (*Room, error) {
	var room *Room
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Print("Failed to genearte uuid\n")
		return room, err
	}

	room = &Room{clients: make(map[*Client]string), name: uuid.String()}
	rooms[uuid.String()] = room
	return room, nil
}
