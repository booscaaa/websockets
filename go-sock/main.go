package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

// Msg é a definição da estrutura de mensagem que iremos passar em ambos os lados da aplicação
type Msg struct {
	Msg   string `json:"msg, omitempty"`
	Room  string `json:"room, omitempty"`
	Red   int    `json:"red, omitempty"`
	Green int    `json:"green, omitempty"`
	Blue  int    `json:"blue, omitempty"`
}

// CustomServer irá ser utilizado para facilitar a alteração do CORS do serviço
type CustomServer struct {
	Server *socketio.Server
}

// ServeHTTP é a funcção que irá substituir o CORS default do serviço
func (s *CustomServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	s.Server.ServeHTTP(w, r)
}

// main é onde a aplicação é executada, é a primeira função que é chamada ao executar o arquivo main.go
func main() {
	ioServer := configureSocketIO()

	wsServer := new(CustomServer)
	wsServer.Server = ioServer

	println("Serviço sendo escutado na porta 5000...")
	http.Handle("/socket.io/", wsServer)
	http.ListenAndServe(":5000", nil)
}

// configureSocketIO é onde definimos os métodos do web socket, como connection por exemplo
func configureSocketIO() *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {

		log.Println("on connection")

		so.On("join", func(msg string) {
			so.Join(msg)
		})

		so.On("message", func(msg string) {
			msgS := Msg{}
			err := json.Unmarshal([]byte(msg), &msgS)
			if err == nil {
				log.Println(msgS.Room)
			}
			log.Println(msg)
			so.BroadcastTo(msgS.Room, "message", msgS)
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
