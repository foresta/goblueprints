package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// clientはチャットを行っている一人のユーザーを表す
type client struct {
	// socketはクライアントのためのwebsocket
	socket *websocket.Conn
	// sendはメッセージが送られるチャネル
	send chan []byte
	// roomはこのクライアントが参加しているチャットルームです
	room *room
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
	c.socket.Close()
}

func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			fmt.Printf("Failed socket.WriteMessage): %s", err)
			return
		}
	}
}
