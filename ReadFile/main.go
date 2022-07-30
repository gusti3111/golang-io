package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	// read strings
	text := strings.NewReader("INI FILE READER\n")
	b := make([]byte, 8)
	_, err := io.ReadFull(text, b)
	if err != nil {
		panic(err)
	}
	// tempfile
	temp, err := ioutil.TempFile("", "temp")
	if err != nil {
		panic(err)
	}
	defer temp.Close()
	fmt.Println(temp.Name())
	// ioutil read
	con, err := ioutil.ReadFile("/tmp/temp122921827")
	if err != nil {
		panic(err)
	}
	fmt.Printf("isi file : %s", con)
	// Write Temp
	if _, err := temp.Write(b); err != nil {
		panic(err)
	}

}
