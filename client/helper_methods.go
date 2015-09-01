package client

import (
	"encoding/json"
)

func (ws *Conn) SendInterface(message interface{}) {
	// Turn the message into json bytes
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return
	}
	// Dump it to clients
	ws.Write(messageBytes)
}

func (ws *Conn) MapInterface(mapObj interface{}) (map[string]interface{}, error) {
	asBytes, err := json.Marshal(mapObj)
	if err != nil {
		return nil, err
	}
	return ws.MapBytes(asBytes)
}

func (ws *Conn) MapBytes(asBytes []byte) (map[string]interface{}, error) {
	var loadValue interface{}
	err := json.Unmarshal(asBytes, &loadValue)
	if err != nil {
		return nil, err
	}
	asMap := loadValue.(map[string]interface{})
	return asMap, nil
}
