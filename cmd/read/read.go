package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	conn, _ := net.Dial("unix", "/tmp/yukino.sock")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, text + "\n")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: "+message)
}