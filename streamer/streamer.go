package streamer

import (
	"log"

	"github.com/gofrs/uuid"
)

type receiveData struct {
	clientID uuid.UUID
	roomID   string
	body     []byte
}

type Streamer struct {
	clientsMap map[uuid.UUID]*client
	roomIDMap  map[string][]uuid.UUID
	listener   chan receiveData
}

func NewStreamer() *Streamer {
	return &Streamer{
		clientsMap: make(map[uuid.UUID]*client),
		roomIDMap:  make(map[string][]uuid.UUID),
		listener:   make(chan receiveData),
	}
}

func (s *Streamer) Listen() {
	for {
		data := <-s.listener

		err := s.handleWebSocket(data)
		if err != nil {
			log.Printf("failed to handle websocket: %v", err)
		}
	}
}

func (s *Streamer) sendToRoom(roomID, msg string) {
	for _, id := range s.roomIDMap[roomID] {
		s.clientsMap[id].sender <- msg
	}
}
