package service

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pdxjohnny/microsocket/client"
)

type Service struct {
	*client.Conn
	// Strings as keys funcitons to call as values
	Methods map[string]func(interface{}, []byte)
	// The struct to call methods
	Caller interface{}
}

type MethodCall struct {
	Method string
}

func NewService() *Service {
	inner := client.NewClient()
	service := Service{
		Conn: inner,
	}
	// Set Recv to MethodMap which will call the correct method
	service.Recv = service.MethodMap
	return &service
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
	fmt.Println(reflect.TypeOf(service.Methods))
	fmt.Println(service.Methods[message.Method])
	service.Methods[message.Method](service.Caller, raw_message)
}
