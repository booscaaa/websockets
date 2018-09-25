package main

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/googollee/go-socket.io"
)

type Msg struct {
	Msg  string `json:"msg, omitempty"`
	Room string `json:"room, omitempty"`
}

//Custom server which basically only contains a socketio variable
//But we need it to enhance it with functions
type customServer struct {
	Server *socketio.Server
}

//Header handling, this is necessary to adjust security and/or header settings in general
//Please keep in mind to adjust that later on in a productive environment!
//Access-Control-Allow-Origin will be set to whoever will call the server
func (s *customServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.Server.ServeHTTP(w, r)
}

func main() {
	//get/configure socket.io websocket for clients
	ioServer := configureSocketIO()

	wsServer := new(customServer)
	wsServer.Server = ioServer

	//HTTP settings
	println("Core Service is listening on port 8081...")
	http.Handle("/socket.io/", wsServer)
	http.ListenAndServe(":5000", nil)
}

func configureSocketIO() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	//Client connects to server
	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")

		so.On("join", func(msg string) {
			so.Join(msg)
		})

		so.On("message", func(msg string) {
			log.Println(msg)
			msgS := Msg{}
			err := json.Unmarshal([]byte(msg), &msgS)
			if err == nil {
					log.Println(msgS.Room)
			}
			so.BroadcastTo(msgS.Room, "message", msgS.Msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}
