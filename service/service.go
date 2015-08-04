package service

import (
	"encoding/json"
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
	// Call the method by name from the pointer to struct service.Caller
	boundMethod := reflect.ValueOf(service.Caller).MethodByName(message.Method)
	// Make sure we have a callable method
	if !boundMethod.IsValid() {
		return
	}
	// Create an argument list for the method
	args := []reflect.Value{reflect.ValueOf(raw_message)}
	// Call the method
	boundMethod.Call(args)
}
