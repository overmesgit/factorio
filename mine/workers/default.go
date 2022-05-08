package workers

import (
	pb "github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine"
)

type DefaultWorker interface {
	GetStorage() Storage
	GetSender() mine.Sender
	GetNextNode() Node
	GetPreviousNode() Node
}

func SendItemFromStore(worker DefaultWorker) error {
	storage := worker.GetStorage()
	Sugar.Infof("Send items. Current store: %v", storage)

	forSend := storage.GetAnyItem()

	if forSend == nil {
		Sugar.Infof("Nothing to send.")
		return nil
	}
	sender := worker.GetSender()
	nextNode := worker.GetNextNode()
	err := sender.SendItem(&pb.Node{Row: nextNode.row, Col: nextNode.col}, forSend)
	if err != nil {
		err := storage.Add(ItemType(forSend.Type))
		if err != nil {
			Sugar.Warnf("Could not stack item back %v %v", forSend, err)
		}
		return err
	}

	return nil
}
