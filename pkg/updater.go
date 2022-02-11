package main

import (
	"time"
	"fmt"
	"sync"
)

func updaterService(wg *sync.WaitGroup) {
	fmt.Println("Starting updater")
	defer wg.Done()
	for {
		mutex.Lock()
		files := make([]FileData, len(data))
		i := 0
		for filepath := range data {
			files[i] = data[filepath]
			i++
		}
		mutex.Unlock()
		currtime := time.Now().Unix()
		for _, file := range files {
			lastupdated := file.LastUpdated
			timeout := file.Timeout
			if currtime > lastupdated + timeout {
				contents, err := refresh(file.Path)
				if err != nil {
					fmt.Println("couldn't update file", err)
					continue
				}
				fmt.Println("updating ", file.Path)
				go updateData(file.Path, contents, file.Timeout)
			}
		}
	}
}
