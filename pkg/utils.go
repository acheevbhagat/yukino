package utils

import (
	"fmt"
	"bufio"
	"os"
	"io"
	"unicode"
)

func read(filepath string) string {
	if contents, ok := data[filepath]; ok {
		return contents
	} else {
		return refresh(filepath)
	}
}

func refresh(filepath string) string {
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