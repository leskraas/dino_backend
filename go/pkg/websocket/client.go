package websocket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/leskraas/dino/pkg/fileManager"
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

func executeYosys(path string) {
	app := "../yosys/yosys"
	// arg0 := []string{"-p", "prep -auto-top -ifx; write_json ../files/out.json", path}
	arg0 := []string{"-p", "prep -auto-top -flatten; write_json ../files/out.json", path}
	yosysCmd := exec.Command(app, arg0...) //, arg1, arg2)

	err := yosysCmd.Run()
	if err != nil {
		// log.Fatalf("yosysCmd.Run() failed with '%s'\n", err)
		log.Printf("yosysCmd.Run() failed with '%s'\n", err)
	}

}

func executeNetlistsvg() {
	app := "node"
	arg := []string{"../netlistsvg/bin/netlistsvg", "../files/out.json", "-o", "../files/out.svg"}
	netlistsvgCmd := exec.Command(app, arg...)

	err := netlistsvgCmd.Run()
	if err != nil {
		// log.Fatalf("netlistsvgCmd.Run() failed with '%s'\n", err)
		log.Printf("netlistsvgCmd.Run() failed with '%s'\n", err)
	}
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
		if d.Type == "verify" {
			fileManager.WriteToFile("../files/ny", d.Content)
			executeYosys("../files/ny.v")
			executeNetlistsvg()
			file, _ := os.Open("../files/out.svg")
			svg, _ := ioutil.ReadAll(file)
			message.Data.Type = "svg"

			message.Data.Content = string(svg)

		}
		c.Pool.Broadcast <- message

		fmt.Printf("Message Received: %+v\n", d.Content)
	}
}
