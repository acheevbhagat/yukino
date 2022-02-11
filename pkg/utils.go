package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"unicode"
	"errors"
	"time"
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

func updateData(filepath string, contents string, timeout int64) {
	mutex.Lock()
	currtime := time.Now().Unix()
	data[filepath] = FileData{
		Path: filepath,
		Contents: contents,
		Timeout: timeout,
		LastUpdated: currtime,
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
