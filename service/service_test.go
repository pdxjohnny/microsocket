package service

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/pdxjohnny/microsocket/config"
	"github.com/pdxjohnny/microsocket/random"
)

type TestService struct {
	*Service
	CorrectResponse chan int
	ShouldBe        string
}

func NewTestService() *TestService {
	inner := NewService()
	service := TestService{Service: inner}
	service.Caller = &service
	return &service
}

type TestServiceMessage struct {
	Method string
	Data   string
}

func (service *TestService) TestServiceMessage(raw_message []byte) {
	// Create a new message struct
	message := new(TestServiceMessage)
	// Parse the message to a json
	json.Unmarshal(raw_message, &message)
	if service.ShouldBe == string(message.Data) {
		service.CorrectResponse <- 1
	}
}

func TestServiceCallMethod(t *testing.T) {
	conf := config.Load()
	correctResponse := make(chan int)
	randString := random.Letters(25)
	service := NewTestService()
	service.ShouldBe = randString
	service.CorrectResponse = correctResponse
	// service.Methods = map[string]func(interface{}, []byte){
	// 	"TestServiceMessage": checkResponse,
	// }
	wsUrl := fmt.Sprintf("http://%s:%s/ws", conf.Host, conf.Port)
	err := service.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	go service.Read()
	log.Println("Waiting for correctResponse", randString)
	checkJson := fmt.Sprintf("{\"data\": \"%s\", \"method\": \"TestServiceMessage\"}", randString)
	service.Write([]byte(checkJson))
	<-correctResponse
	log.Println("Got correctResponse", randString)
}
