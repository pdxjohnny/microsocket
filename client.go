package main

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func Start(url_string string) (ws *websocket.Conn, err error) {
	u, err := url.Parse(url_string)
	if err != nil {
		return nil, err
	}

	rawConn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return nil, err
	}

	wsHeaders := http.Header{
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	wsConn, resp, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
	if err != nil {
		fmt.Errorf("websocket.NewClient Error: %s\nResp:%+v", err, resp)
		return nil, err
	}
	return wsConn, nil
}

func readLoop(conn *websocket.Conn) (err error) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		fmt.Println(string(p))
		// err = conn.WriteMessage(messageType, p)
		// if err != nil {
		// 	return err
		// }
	}
}

func main() {
	ws, err := Start("http://localhost:8080/ws")
	if err != nil {
		fmt.Println(err)
	}
	readLoop(ws)
}
