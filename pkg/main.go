package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

var data map[string]string

func main() {
	data = make(map[string]string)
	l, err := net.Listen("unix", "/tmp/yukino.sock")
	if err != nil {
		println("listen error", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			println("accept error", err)
		}

		go getCommand(conn)
	}
}

func getCommand(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	cmd := strings.Split(message, " ")
	if cmd[0] == "read" {
		filepath := strings.TrimSpace(cmd[1])
		response := read(filepath)
		fmt.Println(response)
		fmt.Fprintf(conn, response + "\n")
	}
	if cmd[0] == "refresh" { 
		filepath := strings.TrimSpace(cmd[1])
		response := read(filepath)
		fmt.Println(response)
		fmt.Fprintf(conn, response + "\n")
	}
}

