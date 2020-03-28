package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)
}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client:" + reply)

		msg := "Received:" + reply
		fmt.Println("Sending to client:" + msg)
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.Handle("/", websocket.Handler(Echo))

	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
