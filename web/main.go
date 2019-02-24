package web

import (
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("upgrade error: %s", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("read error: %s", err)
			break
		}
		log.Printf("Got: %s", string(message))
		// err = c.WriteMessage(mt, message)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	buff, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Printf("error opening file: %s", err)
		return
	}
	w.Write(buff)
}

func Main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/ws", ws)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
