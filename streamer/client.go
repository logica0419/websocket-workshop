package streamer

import (
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type client struct {
	id       uuid.UUID
	roomID   string
	conn     *websocket.Conn
	receiver chan receiveData
	sender   chan string
	closer   chan bool
}

func newClient(roomID string, conn *websocket.Conn, listener chan receiveData) *client {
	return &client{
		id:       uuid.Must(uuid.NewV4()),
		roomID:   roomID,
		conn:     conn,
		receiver: listener,
		sender:   make(chan string),
		closer:   make(chan bool),
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
			payload:     msg,
		}

		c.receiver <- data
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
