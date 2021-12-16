package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"html/template"
)

var (
	conns = map[*websocket.Conn]bool{}
	openConns = make(chan *websocket.Conn)
	closedConns = make(chan *websocket.Conn)
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	conns[conn] = true

	go func() {
		openConns <- conn
	}()

	for {
		select {
		case conn := <-openConns:
			go broadcast(conn)
		case conn := <-closedConns:
			for item := range conns {
				if item == conn {
					delete(conns, item)
					break
				}
			}
		}
	}

}

func broadcast(conn *websocket.Conn) {
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Connection closed:", err)
			break
		}
		fmt.Println("Receiving a message from connection:", string(msg))

		for item := range conns {
			if err := item.WriteMessage(msgType, msg); err != nil {
				log.Println("Channel closed for writing")
				closedConns <- item
			}
		}
	}

	closedConns <- conn
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { 
		template.Must(template.ParseFiles("index.html")).Execute(w, nil)
	})

	http.HandleFunc("/ws", handler)

	http.ListenAndServe(":8000", nil)
}
