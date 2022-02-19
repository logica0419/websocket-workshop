package streamer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type payload struct {
	Method string          `json:"method,omitempty"`
	Args   json.RawMessage `json:"args,omitempty"`
}

type messageArgs struct {
	Message string `json:"message"`
}

func (s *Streamer) handleWebSocket(data receiveData) error {
	var req payload
	err := json.Unmarshal(data.payload, &req)
	if err != nil {
		return err
	}

	switch req.Method {
	case "message":
		var args messageArgs
		err = json.Unmarshal(req.Args, &args)
		if err != nil {
			return err
		}
		s.sendToRoom(data.roomID, args.Message)
	case "time":
		s.sendToRoom(data.roomID, time.Now().Format("01/02 15:04:05"))
	default:
		log.Printf("unknown method: %s", req.Method)
		return fmt.Errorf("unknown method: %s", req.Method)
	}

	return nil
}
