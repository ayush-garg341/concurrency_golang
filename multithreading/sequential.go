package multithreading

import (
	"fmt"
	"math"
	"time"
)

// Counting prime numbers upto 1 million sequentially..

var num int = 1000000
var count int = 0

func CountSeq() {
	start := time.Now()

	for i := 3; i < num; i++ {
		checkPrimeTwo(i)
	}

	elapsed := time.Since(start)

	fmt.Println("Count of prime numbers is ", count+1)
	fmt.Println("Time taken ", elapsed)
}

func checkPrime(n int) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return
		}
	}
	count++
}

func checkPrimeTwo(x int) {
	if x&1 == 0 {
		return
	}
	for i := 3; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	count++
}
