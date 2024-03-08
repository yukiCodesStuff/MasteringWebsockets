package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

// can use an unbuffered channel to handle many requests

func (c *Client) readMessages() {
	defer func() {
		// cleanup connection
		c.manager.removeClient(c)
	}()

	// run forever checking for messages
	for {
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			// can return error is unexpectedly slow
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) { // will throw if reading from a closed connection
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}
