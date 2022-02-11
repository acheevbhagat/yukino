package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"sync"
)

func runCommand(l net.Listener, wg *sync.WaitGroup) {
	fmt.Println("Starting server")
	defer wg.Done()
	for {
		conn, err := l.Accept()
		if err != nil {
			println("accept error", err)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		cmd := strings.Split(message, " ")
		if cmd[0] == "read" {
			runRead(conn, cmd)
		}
		if cmd[0] == "refresh" { 
			runRefresh(conn, cmd)
		}
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
		runRefresh(conn, cmd)
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
		if err != nil {
			timeout = tout
		}
	}
	updateData(filepath, response, int64(timeout))
	fmt.Println(response)
	fmt.Fprintf(conn, response + "\n")
}
