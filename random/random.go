package random

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().Unix())
}

func Letters(length int) string {
	currByte := make([]rune, length)
	for i := range currByte {
		currByte[i] = letters[rand.Intn(len(letters))]
	}
	return string(currByte)
}
