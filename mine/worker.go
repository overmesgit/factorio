package mine

import (
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/localmap"
	"time"
)

func (s *server) RunWorker() {
	go s.DoWork()
	go s.SendItems()
}

func (s *server) DoWork() {
	for {
		time.Sleep(time.Second)
		if MyNode == nil {
			s.logger.Infof("Waiting for my node %v\n", MyNode)
			continue
		}

		s.logger.Infof("Do work %v\n", MyNode)

		switch localmap.Type(MyNode.Type) {
		case localmap.IronMine:
			s.ironMine()
		default:

		}
	}
}

func (s *server) ironMine() {
	mineType := localmap.Iron

	MyItems.Lock()
	defer MyItems.Unlock()

	localStore, ok := MyItems.items[mineType]
	if !ok {
		localStore = &pb.Item{
			Type:  string(mineType),
			Count: 0,
		}
		MyItems.items[mineType] = localStore
	}

	if localStore.Count < 100 {
		localStore.Count++
	}
	s.logger.Infof("IronMine. Dig. LocalStore %v\n", localStore)
}
