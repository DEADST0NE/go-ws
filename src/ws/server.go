package ws

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func WsServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	log.Print("Connect new user")

	if err != nil {
		log.Println("Error WS SERVER connect:", err)
		return
	}

	defer conn.Close()

	client := newClient(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error WS SERVER read message:", err)
			TradeDropSubscriber(&client)
			break
		}

		var message SubMessage
		err = json.Unmarshal(msg, &message)

		if err != nil {
			log.Println("Error WS SERVER parse message:", err)
			SendMessage(conn, "Invalid message")
			continue
		}

		switch message.Ch {
		case "trades":
			TradeHandler(&client, message)
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
