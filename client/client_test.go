package client

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/pdxjohnny/dist-rts/config"
	"github.com/pdxjohnny/dist-rts/server"
)

var randLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(length int) string {
	currByte := make([]rune, length)
	for i := range currByte {
		currByte[i] = randLetters[rand.Intn(len(randLetters))]
	}
	return string(currByte)
}

func checkMessage(should_be string, correctResponse chan int) func(message []byte) {
	return func(message []byte) {
		if should_be == string(message) {
			correctResponse <- 1
		}
	}
}

func TestClientSendRecv(t *testing.T) {
	conf := config.Load()
	go server.Run()
	correctResponse := make(chan int)
	randString := randSeq(50)
	ws := NewClient()
	ws.Recv = checkMessage(randString, correctResponse)
	wsUrl := fmt.Sprintf("http://%s:%s/ws", conf.Host, conf.Port)
	err := ws.Connect(wsUrl)
	if err != nil {
		log.Println(err)
	}
	go ws.Read()
	log.Println("Waiting for correctResponse", randString)
	ws.Write([]byte(randString))
	<-correctResponse
	log.Println("Got correctResponse", randString)
}
