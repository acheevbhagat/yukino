package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"unicode"
	"errors"
)

func readData(filepath string) (FileData, error) {
	mutex.Lock()
	successFlag := false
	err := errors.New("Unknown key")
	val, ok := data[filepath]
	if ok {
		fmt.Println("Key already exists : ", filepath)
		successFlag = true
	} else {
		val = FileData{}	 
	}
	mutex.Unlock()
	if successFlag {
		return val, nil
	} else {
		return val, err
	}
}

func getMinTimeout() (int) {
	mutex.Lock()
	timeout := minTimeout
	mutex.Unlock()
	return timeout
}

func updateData(filepath string, contents string, timeout int) {
	mutex.Lock()
	data[filepath] = FileData{Path: filepath, Contents: contents, Timeout: timeout}
	if timeout < minTimeout {
		minTimeout = timeout
	}
	mutex.Unlock()
}

func read(filepath string) (string, error) {
	data, err := readData(filepath)
	if err != nil {
		return "", errors.New("Unknown key")
	} else { 
		return data.Contents, nil
	}
}

func refresh(filepath string) (string, error) {
	fmt.Println("Reading : ", filepath)
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return "", err
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
	return contents, nil
}