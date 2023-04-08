package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct{}

var instance *single

func getInstance() *single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &single{}
			fmt.Println("New instance created")
		} else {
			fmt.Println("Instance just created")
		}
	} else {
		fmt.Println("Instance already created")
	}

	return instance
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			getInstance()
		}()
	}

	wg.Wait()

}
