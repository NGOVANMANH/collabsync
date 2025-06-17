package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var hub *Hub

func InitHub() {
	hub = NewHub()
	go hub.Run()
}

func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		ID:   r.RemoteAddr,
		Conn: conn,
		Send: make(chan []byte, 256),
		Hub:  hub,
	}

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}
