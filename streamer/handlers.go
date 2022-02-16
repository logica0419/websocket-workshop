package streamer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type body struct {
	Method string                 `json:"method,omitempty"`
	Args   map[string]interface{} `json:"args,omitempty"`
}

func (s *Streamer) handleWebSocket(data receiveData) error {
	var req body
	err := json.Unmarshal(data.body, &req)
	if err != nil {
		return err
	}

	switch req.Method {
	case "message":
		s.sendToRoom(data.roomID, req.Args["message"].(string))
	case "time":
		s.sendToRoom(data.roomID, time.Now().String())
	default:
		log.Printf("unknown method: %s", req.Method)
		return fmt.Errorf("unknown method: %s", req.Method)
	}

	return nil
}
