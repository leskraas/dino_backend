package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int  `json:"type"`
	Data Data `json:"data"`
}

type Data struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		// _, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// fmt.Printf("Message Received: %+v\n", message)
		var d Data
		if err := json.Unmarshal(p, &d); err != nil {
			panic(err)
		}
		message := Message{Type: messageType, Data: d}
		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", d.Content)
	}
}
