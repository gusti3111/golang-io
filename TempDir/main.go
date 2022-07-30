package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	// Read Strings
	t := strings.NewReader("INI ADALAH CARA BUAT TEMP FOLDER\n")
	b, err := ioutil.ReadAll(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
	// make tempfolder
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(dir)
	j := filepath.Join(dir, "tmp")
	if err := ioutil.WriteFile(j, b, 0001); err != nil {
		panic(err)
	}

}
