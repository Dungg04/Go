package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Data one")

	err := ioutil.WriteFile("file.data", data, 0777)

	if err != nil {
		fmt.Println(err)
	}

	dt, err := ioutil.ReadFile("file.data")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dt))

	f, err := os.OpenFile("file.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString("new data"); err != nil {
		panic(err)
	}

	dt, err = ioutil.ReadFile("file.data") 
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(dt))
}