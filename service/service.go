package service

import (
	"encoding/json"
	"fmt"

	"github.com/pdxjohnny/dist-rts/client"
)

type Service struct {
	client.Conn
	// Strings as keys funcitons to call as values
	Methods map[string]Method
}

type MethodCall struct {
	Method string
}

type Method func(*Service, []byte)

func NewService() *Service {
	service := new(Service)
	// Set Recv to MethodMap which will call the correct method
	service.Recv = service.MethodMap
	return service
}

func (service *Service) MethodMap(raw_message []byte) {
	// Create a new message struct
	message := new(MethodCall)
	// Parse the message to a json
	err := json.Unmarshal(raw_message, &message)
	// Make sure there is a method to call and no err
	if err != nil || message.Method == "" {
		return
	}
	fmt.Println("Method", message.Method)
	service.Methods[message.Method](service, raw_message)
}
