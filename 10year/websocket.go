package main

import (
	"fmt"
	//"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// main START OMIT
func echoHandler(ws *websocket.Conn) {
	log.Print("data is recived.")
	//io.Copy(ws, ws)
	buf := make([]byte, 100)
	for {
		ws.Read(buf)
		log.Print(string(buf))
		ws.Write(buf)
	}
}

func main() {
	fmt.Println("please connect to localhost:7778/10year")
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":7778", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

// main END OMIT
