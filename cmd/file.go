package main

import (
	"learning-web-socket/internal/handlers"
	"log"
	"net/http"
)

func main() {
	routes := routes()
	log.Println("Staring channel listener...")

	//serve static folder
	go handlers.ListenToWSchannel()
	_ = http.ListenAndServe(":8080", routes)

}
