package multithreading

// Counting prime numbers upto 1 million in threaded fair way

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

var start int32 = 2
var numfair int32 = 10000000
var countfair int32 = 0

func CountSeqFair() {
	begin := time.Now()

	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go threadedPrimeCountFair(i, &wg)
	}
	wg.Wait()
	elapsed := time.Since(begin)

	fmt.Println("Count of prime numbers is ", countUnfair+1, " time taken : ", elapsed)
}

func threadedPrimeCountFair(i int, wg *sync.WaitGroup) {
	begin := time.Now()
	for {
		x := atomic.AddInt32(&start, 1)
		if x > numfair {
			break
		}
		checkPrimeFair(x)
	}
	fmt.Printf("Thread %d completed in %s\n", i, time.Since(begin))
	wg.Done()
}

func checkPrimeFair(num int32) {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if int(num)%i == 0 {
			return
		}
	}
	atomic.AddInt32(&countUnfair, 1)
}
