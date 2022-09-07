package streamer

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

func (s *Streamer) ConnectWS(c echo.Context) error {
	roomID := c.QueryParam("room")
	if roomID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "roomID is required")
	}

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	client := newClient(roomID, conn, s.receiver)

	s.clients[client.id] = client

	go client.listen()
	go client.send()

	<-client.closer

	delete(s.clients, client.id)

	return c.NoContent(http.StatusOK)
}
