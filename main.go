package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"github.com/eye1994/game-server/ws"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", ws.HandleWsConnection)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
