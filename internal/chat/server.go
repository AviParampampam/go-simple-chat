package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

// Server ..
type Server struct {
	Port int
	Conn net.Conn
}

// CreateServer ..
func CreateServer(port int) (Server, error) {
	var s Server
	s.Port = port
	return s, nil
}

// Listen ..
func (s *Server) Listen() {
	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(s.Port))

	// Открываем порт
	conn, _ := ln.Accept()
	s.Conn = conn

	// Запускаем цикл
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("\nСобеседник: %s\n", string(message))
	}
}

// SendMessage ..
func (s *Server) SendMessage() {
	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')

	s.Conn.Write([]byte(msg + "\n"))
}
