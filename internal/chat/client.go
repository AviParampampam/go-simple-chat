package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Client ..
type Client struct {
	Conn net.Conn
}

// CreateClient ..
func CreateClient() Client {
	var c Client
	return c
}

// Connect ..
func (c *Client) Connect(addr string) {
	conn, _ := net.Dial("tcp", addr)
	c.Conn = conn
}

// ListenOnMessage ..
func (c *Client) ListenOnMessage() {
	for {
		message, _ := bufio.NewReader(c.Conn).ReadString('\n')
		fmt.Printf("\nСобеседник: %s\n", message)
	}
}

// SendMessage ...
func (c *Client) SendMessage() {

	reader := bufio.NewReader(os.Stdin)
	msg, _ := reader.ReadString('\n')

	fmt.Fprintf(c.Conn, msg+"\n")
}
