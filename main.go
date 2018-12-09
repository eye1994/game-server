package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/eye1994/game-server/ws"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func ping(w http.ResponseWriter, r *http.Request) {
		 fmt.Fprintf(w, "pong")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ws", ws.HandleWsConnection)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
