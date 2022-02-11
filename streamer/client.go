package streamer

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id       uuid.UUID
	roomID   string
	conn     *websocket.Conn
	listener *chan receiveData
	sender   chan string
	closer   chan bool
}

func newClient(roomID string, conn *websocket.Conn, listener *chan receiveData) *client {
	return &client{
		id:       uuid.Must(uuid.NewV4()),
		roomID:   roomID,
		conn:     conn,
		listener: listener,
		sender:   make(chan string),
	}
}

func (c *client) listen() {
	for {
		msgType, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.closer <- true
			return
		}
		if msgType != websocket.TextMessage {
			continue
		}

		data := receiveData{
			clientID: c.id,
			roomID:   c.roomID,
			data:     string(msg),
		}

		*c.listener <- data
	}
}

func (c *client) send() {
	for {
		msg := <-c.sender

		err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			c.closer <- true
			return
		}
	}
}
