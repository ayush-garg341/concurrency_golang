package main

import (
	//race "concurrency_go/race"
	//lock "concurrency_go/locking"
	mt "concurrency_go/multithreading"
)

func main() {
	//race.RaceConditions()
	//race.RacyChannel()
	//mt.CountSeq()
	//lock.Count()
	mt.CountSeqUnfair()
	//mt.CountSeq()
}
