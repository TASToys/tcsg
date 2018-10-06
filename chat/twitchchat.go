package chat

import (
	"log"
	"os"
	"os/signal"
	"strings"

	nats "github.com/nats-io/go-nats"
	"github.com/sacOO7/gowebsocket"
)

func Connect() {
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Panic(err)
	}
	c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		log.Panic(err)
	}

	socket := gowebsocket.New("wss://irc-ws.chat.twitch.tv")

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Received connect error - ", err)
	}

	socket.OnConnected = func(socket gowebsocket.Socket) {
		socket.SendText("CAP REQ :twitch.tv/tags twitch.tv/commands twitch.tv/membership")
		socket.SendText("NICK justinfan64546468")
		joinChannel(socket, "twitchpresents")
		// joinChannel(socket, "twitchpresents2")
		// joinChannel(socket, "scoga")
		joinChannel(socket, "tinahacks")
		joinChannel(socket, "illyohs")
		joinChannel(socket, "carlsagan42")
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		// log.Println("Received message - " + message)

		if strings.Contains(message, "PRIVMSG") {

			string2Prv(c, message)
		}
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received ping - " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Received pong - " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	for {
		select {
		case <-interupt:
			log.Println("interupt")
			socket.Close()
			return
		}
	}
}

func joinChannel(socket gowebsocket.Socket, channel string) {
	socket.SendText("JOIN #" + channel)
}
