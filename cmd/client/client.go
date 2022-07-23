package main

import (
	"fmt"
	"runtime"
)

func main() {

	_ = make([]int64, 1_000_000)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	//runtime.Stat

	//fmt.Printf("%+v", m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
