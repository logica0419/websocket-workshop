package streamer

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func (s *Streamer) ConnectWS(c echo.Context) error {
	roomID := c.QueryParam("room")
	if roomID == "" {
		return c.String(http.StatusBadRequest, "roomID is required")
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("failed to upgrade connection: %v", err))
	}

	client := newClient(roomID, conn, &s.listener)

	s.clientsMap[client.id] = client
	s.roomIDMap[roomID] = append(s.roomIDMap[roomID], client.id)

	go client.listen()
	go client.send()

	<-client.closer

	delete(s.clientsMap, client.id)
	for i, id := range s.roomIDMap[roomID] {
		if id == client.id {
			s.roomIDMap[roomID] = append(s.roomIDMap[roomID][:i], s.roomIDMap[roomID][i+1:]...)
			break
		}
	}

	return c.NoContent(http.StatusOK)
}
