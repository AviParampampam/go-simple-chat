package main

import "github.com/AviParampampam/go-chat/internal/chat"

func main() {
	client := chat.CreateClient()
	client.Connect("127.0.0.1:22222")
	go client.ListenOnMessage()
	for {
		client.SendMessage()
	}
}
