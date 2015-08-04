package service

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/pdxjohnny/microsocket/config"
	"github.com/pdxjohnny/microsocket/random"
	"github.com/pdxjohnny/microsocket/server"
)

type TestService struct {
	Method string
	Data   string
}

func checkMethod(should_be string, correctResponse chan int) func(service *Service, raw_message []byte) {
	return func(service *Service, raw_message []byte) {
		// Create a new message struct
		message := new(TestService)
		// Parse the message to a json
		json.Unmarshal(raw_message, &message)
		if should_be == string(message.Data) {
			correctResponse <- 1
		}
	}
}

func TestServiceCallMethod(t *testing.T) {
	conf := config.Load()
	go server.Run()
	correctResponse := make(chan int)
	randString := random.Letters(50)
	service := NewService()
	service.Methods = map[string]Method{
		"TestService": checkMethod(randString, correctResponse),
	}
	wsUrl := fmt.Sprintf("http://%s:%s/ws", conf.Host, conf.Port)
	err := service.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	go service.Read()
	log.Println("Waiting for correctResponse", randString)
	checkJson := fmt.Sprintf("{\"data\": \"%s\", \"method\": \"TestService\"}", randString)
	service.Write([]byte(checkJson))
	<-correctResponse
	log.Println("Got correctResponse", randString)
}
