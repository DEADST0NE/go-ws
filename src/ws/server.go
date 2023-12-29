package ws

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func WsServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3004", nil))
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	log.Info("Client Connected")

	if err != nil {
		log.Println("Error WS SERVER connect:", err)
		return
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	defer conn.Close()
	client := newClient(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Info("Error WS SERVER read message:", err)
			TradeDropSubscriber(&client)
			break
		}

		if string(msg) == "ping" {
			SendMessage(conn, "pong")
			continue
		}
		if string(msg) == "pong" {
			SendMessage(conn, "ping")
		}

		var message SubMessage
		err = json.Unmarshal(msg, &message)

		if err != nil {
			log.Info("Error WS SERVER parse message:", err)
			SendMessage(conn, "Invalid message")
			continue
		}

		switch message.Ch {
		case "trades":
			TradeHandler(&client, message)
		case "rsi":
			RsiHandler(&client, message)
		default:
			SendMessage(conn, "Chanal not implemented")
		}
	}
}

func SendMessage(conn *websocket.Conn, message string) {
	var err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("Error WS SERVER send valid message:", err)
	}
}

func newClient(conn *websocket.Conn) Client {
	return Client{
		ID:   uuid.New().String(),
		Conn: conn,
	}
}
