package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conns = map[*websocket.Conn]bool{}
)

type Chat struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	conns[conn] = true

	go func(conn *websocket.Conn) {
		for {
			var c Chat
			err := conn.ReadJSON(&c)
			if err != nil {
				fmt.Println("could not read from connection: deleting connection", err)
				delete(conns, conn)
				return
			}

			if c.Type == "new user" {
				fmt.Printf("new user name: %v joined\n", c.Name)
			} else if c.Type == "new message" {
				fmt.Printf("Received message: %v\n", c.Msg)
			}

			for eachConn := range conns {
				if err := eachConn.WriteJSON(c); err != nil {
					fmt.Println("could not write to connection: deleting connection", err)
					delete(conns, eachConn)
					continue
				}
			}
		}
	}(conn)
}

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "home.html")
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/newchat", func(w http.ResponseWriter, r *http.Request){
			groupId := uuid.NewV4()
			http.Redirect(w, r, "/group?id="+groupId, http.StatusSeeOther)
		
	})

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "home.html")
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/chat", ChatHandler)

	fmt.Println("visit localhost:8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
