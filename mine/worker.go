package mine

import (
	"github.com/overmesgit/factorio/nodemap"
	"time"
)

func (s *server) RunWorker() {
	go s.DoWork()
	go s.SendItems()
}

func (s *server) DoWork() {
	var err error
	for {
		time.Sleep(time.Second)
		sugar.Infof("Do work %v\n", MyNode)
		err = nil

		switch nodemap.Type(MyNode.Type) {
		case nodemap.IronMine:
			err = s.ironMine()
		default:

		}

		sugar.Infof("After work. Err %v LocalStore %v", err, MyStorage.GetItemCount())

	}
}

func (s *server) ironMine() error {
	return MyStorage.Add(nodemap.Iron)
}
