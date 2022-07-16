package basic

import (
	"github.com/overmesgit/factorio/grpc"
	"github.com/overmesgit/factorio/mine"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers"
)

type Sender struct {
	*Storage
	mine.Sender
	nextNode workers.Node
}

func NewSender(storage *Storage, sender mine.Sender, nextNode workers.Node) Sender {
	return Sender{Storage: storage, Sender: sender, nextNode: nextNode}
}

func (s Sender) SendItemFromStore() error {
	sugar.Sugar.Infof("Send items. Current store: %v", s.Storage)
	forSend, err := s.Storage.GetAnyItem()

	if err != nil {
		sugar.Sugar.Infof("Nothing to send: " + err.Error())
		return err
	}

	err = s.Sender.SendItem(
		&grpc.Node{Row: s.nextNode.Row, Col: s.nextNode.Col}, &grpc.Item{Type: string(forSend)},
	)
	if err != nil {
		err := s.Storage.Add(forSend)
		if err != nil {
			sugar.Sugar.Warnf("Could not stack item back %v %v", forSend, err)
		}
		return err
	}

	return nil
}
