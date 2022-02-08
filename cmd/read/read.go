package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, _ := net.Dial("unix", "/tmp/yukino.sock")
	// what to send?
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send: ")
	text, _ := reader.ReadString('\n')
	// send to server
	fmt.Fprintf(conn, text + "\n")
	// wait for reply
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: "+message)
}