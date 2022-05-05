package mine

import (
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"time"
)

func (s *server) RunWorker() {
	go s.DoWork()
}

func (s *server) DoWork() {
	for {
		time.Sleep(time.Second)
		s.logger.Infof("Do work for %v\n", MyType)
		switch MyType {
		case localmap.IronMine:
			s.ironMine()
		default:
			s.logger.Warnf("Waiting for my type %v\n", MyType)

		}
	}
}

func (s *server) ironMine() {
	mineType := localmap.Iron

	localStore, ok := MyItems[mineType]
	if !ok {
		localStore = &pb.Item{
			Type:  string(mineType),
			Count: 0,
		}
		MyItems[mineType] = localStore
	}

	s.logger.Infof("LocalStore %v\n", localStore)
	if localStore.Count < 100 {
		localStore.Count++
	}
}
