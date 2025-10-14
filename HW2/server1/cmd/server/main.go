package main

import (
	"net/http"
	"server1/internal/config"
	"server1/internal/handlers"
)

func main() {
	http.HandleFunc("/version", handlers.Version)
	http.HandleFunc("/decode", handlers.Decode)
	http.HandleFunc("/hard-op", handlers.HardOp)
	err := http.ListenAndServe(config.GetPort(), nil)
	if err != nil {
		return
	}
}
