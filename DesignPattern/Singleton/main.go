package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type signle struct {}

var signleInstance *signle 

func getInstance() *signle{
	if signleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if signleInstance == nil {
			fmt.Println("Creating single instance now")
			signleInstance = &signle{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return signleInstance
}

func main() {
	for i:=0; i<30; i++{
		go getInstance()
	}

	fmt.Scanln()
}