package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var num int32 = 0

func increment() {
	num += 1
}

func Count() {
	start := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go threadSafeCount(&wg, &num)
		//go threadCount(&wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Total count without locking : ", num, " time taken ", elapsed)
}

func threadCount(wg *sync.WaitGroup) {
	for i := 0; i < 1000000; i++ {
		num += 1
	}
	wg.Done()
}

func threadSafeCount(wg *sync.WaitGroup, num *int32) {
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(num, 1)
	}
	wg.Done()
}
