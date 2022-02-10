package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"os"
	"io"
	"unicode"
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
}

func read(filepath string) string {
	if contents, ok := data[filepath]; ok {
		return contents
	} else {
		fmt.Println("Reading : ", filepath)
		jsonFile, err := os.Open(filepath)
		if err != nil {
			fmt.Println("Fatal error : ", err)
			return ""
		}
		defer jsonFile.Close()
		r := bufio.NewReader(jsonFile)
		contents := ""
		for {
			c, _, err := r.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				} else {
					fmt.Println("Fatal error : ", err)
				}
			} else {
				// Need to do validation as we process each character
				// to make sure we're reading valid json
				if !unicode.IsSpace(c) {
					contents = contents + string(c)
				}
			}
		}
		data[filepath] = contents
		return contents
	}
}