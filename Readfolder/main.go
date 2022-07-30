package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	// Read Strings
	text := strings.NewReader("Ini Cara Read Folder")
	b := make([]byte, 13)
	_, err := io.ReadFull(text, b)
	if err != nil {
		panic(err)
	}

	// tempfile
	tmp, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(tmp.Name())
	defer tmp.Close()

	// Read Folder
	f, err := ioutil.ReadDir("/tmp")
	if err != nil {
		panic(err)
	}
	for _, file := range f {
		fmt.Println(file.Name())

	}

	// Write
	if _, err := tmp.Write(b); err != nil {
		panic(err)
	}

}
