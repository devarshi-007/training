package socketdemo

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// let ws = new WebSocket("ws://localhost:5000/websocket")
// ws.send(JSON.stringify({data: "nothing mean", type: "not null"}))
// ws.onmessage = function(e) {console.log(e.data)}
// ws.close()

func demoSocket(ws *websocket.Conn) {

	done := make(chan struct{})

	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				fmt.Println(err)
				break
			}
			fmt.Println("received message", msg.Data)
		}
		close(done)
	}(ws)

loop:
	for {
		select {
		case <-done:
			fmt.Println("connection closed")
			break loop
		default:
			arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
			rand.Shuffle(len(arr), func(i, j int) {
				arr[i], arr[j] = arr[j], arr[i]
			})
			if err := websocket.JSON.Send(ws, arr); err != nil {
				log.Println(err)
				break
			}
			time.Sleep(5 * time.Second)

		}
	}

	fmt.Println("closing the web socket")
	defer ws.Close()
}

// listeners
func StartListening() {
	http.Handle("/websocket", websocket.Handler(demoSocket))
	http.ListenAndServe(":5000", nil)
}
