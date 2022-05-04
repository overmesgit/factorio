package mine

import (
	"time"
)

func (s *server) RunWorker() {
	go s.DoWork()
}

func (s *server) DoWork() {
	for {
		time.Sleep(time.Second)
		if MyType != "" {
			s.logger.Infof("Do some work %v\n", MyType)
		} else {
			s.logger.Infof("Waiting for mytype %v\n", MyType)
		}
	}
}
