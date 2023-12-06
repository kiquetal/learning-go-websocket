package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections
	},
}

func Home(w http.ResponseWriter, r *http.Request) {

	renderPages(w, "home.jet", nil)

}

type WSJonResponse struct {
	Message     string `json:"message"`
	Action      string `json:"action"`
	MessageType string `json:"message_type"`
}
type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action  string              `json:"action"`
	Message string              `json:"message"`
	Conn    WebSocketConnection `json:"-"`
}

var wsChan = make(chan WsPayload)
var clients = make(map[WebSocketConnection]bool)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	con, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Server >> Client Successfully Connected...")
	var response WSJonResponse
	response.Message = `<em><small>Connected to <b>server</b>!!</small></em>`
	err = con.WriteJSON(response)
	conn := WebSocketConnection{con}
	clients[conn] = true

	if err != nil {
		panic(err)
	}

	go ListerForWS(&conn)

}

func ListerForWS(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error >> ", r)
		}
	}()
	var payload WsPayload
	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			log.Println(err)
			delete(clients, *conn)
			break
		}
		payload.Conn = *conn
		wsChan <- payload
	}
}

func ListenToWSchannel() {
	var response WSJonResponse
	for {
		e := <-wsChan
		response.Action = "gOT HERE"
		response.Message = fmt.Sprintf("Message: %s", e.Message)

		broadcastToAll(response)
	}
}

func broadcastToAll(response WSJonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println(err)
			client.Close()
			delete(clients, client)
		}
	}
}

func reader(conn *websocket.Conn) {
	// create the connection for the websocket

	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		// print out that message for clarity
		println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			panic(err)
		}
	}
}

func renderPages(w http.ResponseWriter, tmpl string, data jet.VarMap) {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
