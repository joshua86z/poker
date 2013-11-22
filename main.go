package main

import (
	"code.google.com/p/go.net/websocket"
	//	"fmt"
	//	"io"
	"github.com/fhbzyc/poker/models/routine"
	"log"
	"net/http"

//	"time"

//	"strconv"
)

func ChatWith(ws *websocket.Conn) {

	routine.WsList[routine.WsListNum] = ws
	routine.WsListNum++

	var err error

	//	t := time.Now().Unix() - 2

	for {

		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			break
		} else {
			//			fmt.Println(reply)
			if reply == "" {
				continue
			}
			go routine.Run(reply, ws)

			//				continue
			//			}
			//			t = time.Now().Unix() - t

		}
	}
}

func main() {
	//
	go http.Handle("/", websocket.Handler(ChatWith))
	http.HandleFunc("/chat", Client)

	//	fmt.Println("listen on port 8001")
	//	fmt.Println("visit http://127.0.0.1:8001/chat with web browser(recommend: chrome)")

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Client(w http.ResponseWriter, r *http.Request) {
	index.Client(w, r)
}
