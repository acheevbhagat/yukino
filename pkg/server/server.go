package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"encoding/json"
	"os"
	"io/ioutil"
)

func main() {
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
		filepath := cmd[1]
		jsonFile, _ := os.Open(filepath)
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)
		fmt.Println(result)
	}
}