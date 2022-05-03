package mine

import (
	"log"
	"time"
)

func RunWorker() {
	go DoWork()
}

func DoWork() {
	for {
		time.Sleep(time.Second)
		if MyType != "" {
			log.Printf("Do some work %v\n", MyType)
		} else {
			log.Printf("Waiting for mytype %v\n", MyType)
		}
	}
}
