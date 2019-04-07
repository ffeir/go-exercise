package main

import (
	"sync"
	"time"
)

var m sync.Mutex

func testWithM(s string) {
	m.Lock()
	defer m.Unlock()

	for i := 0; i < 5; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		testWithM("read")
	}()

	go func() {
		defer wg.Done()
		testWithM("write")
	}()

	wg.Wait()
}
