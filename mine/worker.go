package mine

import (
	"errors"
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
			err = MyStorage.Add(nodemap.Iron)
		case nodemap.CoalMine:
			err = MyStorage.Add(nodemap.Coal)
		case nodemap.Furnace:
			err = produceFurnace()
		case nodemap.Manipulator:
			neededItem, err := s.askForNeedItem(s.getNextNode())
			if err == nil {
				item, err := s.askForItem(s.getPrevNode(), nodemap.ItemType(neededItem.Type), false)
				if err == nil {
					err = s.sendItem(s.getNextNode(), item)
				}
			}
		default:

		}

		sugar.Infof("After work. Err %v LocalStore %v", err, MyStorage.GetItemCount())

	}
}

func produceFurnace() error {
	if MyStorage.GetCount(nodemap.Iron) == 0 || MyStorage.GetCount(nodemap.Coal) == 0 {
		return errors.New("not enough materials")
	}

	if MyStorage.isFull(nodemap.IronPlate) {
		return errors.New("storage is full")
	}

	MyStorage.Get(nodemap.Iron)
	MyStorage.Get(nodemap.Coal)
	return MyStorage.Add(nodemap.IronPlate)
}
