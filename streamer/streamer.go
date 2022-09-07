package streamer

import (
	"log"

	"github.com/gofrs/uuid"
)

type receiveData struct {
	clientID uuid.UUID
	roomID   string
	payload  []byte
}

type Streamer struct {
	clients  map[uuid.UUID]*client
	receiver chan receiveData
}

func NewStreamer() *Streamer {
	return &Streamer{
		clients:  make(map[uuid.UUID]*client),
		receiver: make(chan receiveData),
	}
}

func (s *Streamer) Listen() {
	for {
		data := <-s.receiver

		go func() {
			err := s.handleWebSocket(data)
			if err != nil {
				log.Printf("failed to handle websocket: %v", err)
			}
		}()
	}
}

func (s *Streamer) sendToRoom(roomID, msg string) {
	for _, c := range s.clients {
		if c.roomID == roomID {
			c.sender <- msg
		}
	}
}
