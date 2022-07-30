package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// tempfile

	text := []byte("HAI INI ADALAH TEMPFILE")
	temp, err := ioutil.TempFile("", "temp")
	if err != nil {
		panic("errr")
	}
	fmt.Println(temp.Name())

	// write
	if _, err := temp.Write(text); err != nil {
		panic(err)
	}
	defer temp.Close()
}
