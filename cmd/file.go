package main

import "net/http"

func main() {
	routes := routes()
	_ = http.ListenAndServe(":8080", routes)

}
