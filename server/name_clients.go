package server

import (
	"fmt"
	"hash/fnv"

	"github.com/pdxjohnny/easysock"
)

var Hub = easysock.Hub

func hash(s interface{}) string {
	asString := fmt.Sprintf("%s", s)
	h := fnv.New32a()
	h.Write([]byte(asString))
	return fmt.Sprint(h.Sum32())
}

func init() {
	Hub.OnConnect = func (Hub *easysock.WebSocketManger, Conn *easysock.Connection)  {
		nameHash := hash(Conn)
		Conn.Data = nameHash
		sendName := fmt.Sprintf(`{"Method": "MicroSocketName", "Name": "%s"}`, nameHash)
		Conn.Send <- []byte(sendName)
	}

	Hub.OnClose = func (Hub *easysock.WebSocketManger, Conn *easysock.Connection)  {
		nameHash := Conn.Data.(string)
		sendClosed := fmt.Sprintf(`{"Method": "Closed", "Name": "%s"}`, nameHash)
		Hub.Broadcast <- []byte(sendClosed)
	}
}
