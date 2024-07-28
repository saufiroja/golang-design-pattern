package main

import (
	"fmt"
	"sync"
)

func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}

	fmt.Scanln()
}

type single struct {
}

var instance *single
var lock = &sync.Mutex{}

func getInstance() *single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			fmt.Println("Creating Single Instance Now")
			instance = &single{}
		} else {
			fmt.Println("Single Instance Already Created")
		}
	} else {
		fmt.Println("Single Instance Already Created")
	}

	return instance
}
