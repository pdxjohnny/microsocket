package service

import (
  "fmt"
  "log"
  "encoding/json"

  "github.com/pdxjohnny/dist-rts/client"
)

type Service struct {
	client.Conn
  Methods map[string]interface{}
}

type MethodCall struct {
    Method string
}

func NewService() *Service {
  service := new(Service)
  service.Recv = MethodMap
  return service
}

func MethodMap(raw_message []byte) {
  message := new(MethodCall)
  err := json.Unmarshal(raw_message, &message)
  if err != nil {
    log.Println(err)
  }
  fmt.Println("Method", message.Method)
}
