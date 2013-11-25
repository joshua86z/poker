package main

import (
	"code.google.com/p/go.net/websocket"
	//	"fmt"
	//	"io"
	"github.com/fhbzyc/poker/models"
	"log"
	"net/http"

//	"time"

//	"strconv"
)

func ChatWith(ws *websocket.Conn) {

	if models.WsListNum >= 9 {
		return
	}

	models.WsList = append(models.WsList, ws)
	models.WsListNum++

	if models.WsListNum == 2 {
		go models.Play()
	}

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
			go models.Run(reply, ws)

			//				continue
			//			}
			//			t = time.Now().Unix() - t

		}
	}
}

func main() {
	//
	http.Handle("/", websocket.Handler(ChatWith))
	//	http.HandleFunc("/chat", Client)

	//	fmt.Println("listen on port 8001")
	//	fmt.Println("visit http://127.0.0.1:8001/chat with web browser(recommend: chrome)")

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
