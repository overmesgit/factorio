package workers

import (
	"errors"
	"github.com/overmesgit/factorio/mine"
	"github.com/overmesgit/factorio/mine/sugar"
	"time"
)

type ManipulatorNode struct {
	nextNode, prevNode Node
	sender             mine.Sender
}

var _ WorkerNode = ManipulatorNode{}

func NewManipulator(
	nextNode, prevNode Node,
) *ManipulatorNode {
	return &ManipulatorNode{nextNode: nextNode, prevNode: prevNode, sender: mine.NewSender()}
}

func (n ManipulatorNode) GetNeededResource() (ItemType, error) {
	return NoItem, errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) GetResourceForSend() (ItemType, error) {
	return NoItem, errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) ReceiveResource(itemType ItemType) error {
	return errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) StartWorker() {
	go func() {
		n.transitResource()
	}()
}

func (n ManipulatorNode) transitResource() {
	neededItem, err := n.sender.AskForNeedItem(n.nextNode)
	if err == nil {
		item, err := n.sender.AskForItem(n.prevNode, neededItem)

		if err != nil {
			sugar.Sugar.Infof("Error while asking %v for item %v", n.prevNode, err)
			return
		}
		//if store {
		//err := MyStorage.Add(nodemap.ItemType(r.Type))
		//}

		err = n.sender.SendItem(n.nextNode, item)
		if err != nil {
			sugar.Sugar.Infof("Error while sending %v item %v", n.nextNode, err)
			return
		}
	}
	time.Sleep(time.Second)
}
