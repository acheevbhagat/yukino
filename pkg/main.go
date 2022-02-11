package main

import (
	"net"
	"sync"
)

var mutex sync.Mutex
var data map[string]FileData

func main() {
	var wg sync.WaitGroup
	data = make(map[string]FileData)
	listener, err := net.Listen("unix", "/tmp/yukino.sock")
	if err != nil {
		println("listen error", err)
		return
	}
	wg.Add(2)
	go updaterService(&wg)
	go runCommand(listener, &wg)
	wg.Wait()
}
