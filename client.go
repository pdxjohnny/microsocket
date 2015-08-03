package main

import (
  "fmt"
  "net"
  "net/url"
  "net/http"

  "github.com/gorilla/websocket"
)

func Start(url_string string) interface{} {
  u, err := url.Parse(url_string)
  if err != nil {
      return err
  }

  rawConn, err := net.Dial("tcp", u.Host)
  if err != nil {
      return err
  }

  wsHeaders := http.Header{
      // "Origin":                   {url},
      // your milage may differ
      "Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
  }

  // wsConn, resp, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
  _, resp, err := websocket.NewClient(rawConn, u, wsHeaders, 1024, 1024)
  if err != nil {
      return fmt.Errorf("websocket.NewClient Error: %s\nResp:%+v", err, resp)
  }
  fmt.Println("Did it!!")
  return nil
}

func main() {
  err := Start("http://localhost:8080/ws")
  if err != nil {
    fmt.Println(err)
  }
}
