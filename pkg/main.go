package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"sync"
	"math"
)

var mutex sync.Mutex
var data map[string]FileData
var minTimeout int

func main() {
	data = make(map[string]FileData)
	minTimeout = math.MaxUint32
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

		go runCommand(conn)
	}
}

func runCommand(conn net.Conn) {
	message, _ := bufio.NewReader(conn).ReadString('\n')
	cmd := strings.Split(message, " ")
	if cmd[0] == "read" {
		runRead(conn, cmd)
	}
	if cmd[0] == "refresh" { 
		runRefresh(conn, cmd)
	}
}

func runRead(conn net.Conn, cmd []string) {
	filepath := strings.TrimSpace(cmd[1])
	response, err := read(filepath)
	refreshFlag := false
	if err != nil {
		refreshFlag = true
	}
	if refreshFlag {
		response, err = refresh(filepath)
		if err != nil {
			fmt.Println("Fatal error : ", err)
			return
		}
	} else {
		fmt.Println(response)
		fmt.Fprintf(conn, response + "\n")
	}
}

func runRefresh(conn net.Conn, cmd []string) {
	filepath := strings.TrimSpace(cmd[1])
	response, err := refresh(filepath)
	if err != nil {
		fmt.Println("Fatal error : ", err)
		return
	}
	timeout := 20
	if len(cmd) > 2 {
		tout, err := strconv.Atoi(strings.TrimSpace(cmd[2]))
		if err == nil {
			timeout = tout
		}
	}
	updateData(filepath, response, timeout)
	fmt.Println(response)
	fmt.Fprintf(conn, response + "\n")
}
