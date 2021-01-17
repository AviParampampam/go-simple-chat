package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AviParampampam/go-chat/internal/chat"
	"github.com/BurntSushi/toml"
)

func main() {
	var conf chat.ServerConfig
	tomlData, err := ioutil.ReadFile("configs/server.toml")
	check(err)
	_, err = toml.Decode(string(tomlData), &conf)
	check(err)

	s, _ := chat.CreateServer(conf.Port)
	go s.Listen()
	for {
		s.SendMessage()
	}
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
