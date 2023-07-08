package multithreading

// Counting prime numbers upto 1 million threaded unfair.

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var numUnfair int = 1000000
var countUnfair int32 = 0

func CountSeqUnfair() {
	begin := time.Now()

	wg := sync.WaitGroup{}
	start := 3
	end := 100000
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		end = i * 100000
		go threadedPrimeCount(&wg, start, end)
		start = end + 1
	}
	wg.Wait()
	elapsed := time.Since(begin)

	fmt.Println("Final end value is ", end)
	fmt.Println("Count of prime numbers is ", countUnfair+1, " time taken : ", elapsed)
}

func threadedPrimeCount(wg *sync.WaitGroup, start int, end int) {
	for i := start; i <= end; i++ {
		checkPrimeUnfair(i)
	}
	wg.Done()
}

func checkPrimeUnfair(num int) {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return
		}
	}
	atomic.AddInt32(&countUnfair, 1)
}
