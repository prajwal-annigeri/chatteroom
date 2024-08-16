package ws

import (
	"fmt"
)

type Room struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Clients  map[string]*Client `json:"clients"`
}
type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			if _, ok := hub.Rooms[client.RoomID]; ok {
				r := hub.Rooms[client.RoomID]
				if r.Password == client.Password {
					if _, ok := r.Clients[client.ID]; !ok {
						r.Clients[client.ID] = client
					}
				}
			}
		case client := <-hub.Unregister:
			if _, ok := hub.Rooms[client.RoomID]; ok {
				r := hub.Rooms[client.RoomID]
				if _, ok := r.Clients[client.ID]; ok {
					if len(r.Clients) != 0 {
						hub.Broadcast <- &Message{
							Content:  fmt.Sprintf("%s has left the room", client.Username),
							RoomId:   client.RoomID,
							Username: client.Username,
						}
					}
					delete(r.Clients, client.ID)
					close(client.Message)
				}
			}
		case m := <-hub.Broadcast:
			if _, ok := hub.Rooms[m.RoomId]; ok {
				for _, client := range hub.Rooms[m.RoomId].Clients {
					client.Message <- m
				}
			}
		}
	}
}
