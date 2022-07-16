package workers

import (
	"errors"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type ManipulatorNode struct {
	nextNode, prevNode basic.Node
	sender             basic.Sender
}

var _ basic.WorkerNode = ManipulatorNode{}

func NewManipulator(
	nextNode, prevNode basic.Node, sender basic.Sender,
) *ManipulatorNode {
	return &ManipulatorNode{nextNode: nextNode, prevNode: prevNode, sender: sender}
}

func (n ManipulatorNode) GetItemCount() []basic.ItemCounter {
	return nil
}

func (n ManipulatorNode) GetNeededResource() (basic.ItemType, error) {
	return basic.NoItem, errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) GetResourceForSend(basic.ItemType) (basic.ItemType, error) {
	return basic.NoItem, errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) ReceiveResource(itemType basic.ItemType) error {
	return errors.New("i'm an manipulator dumb dumb")
}

func (n ManipulatorNode) StartWorker() {
	go func() {
		for {
			n.transitResource()
			time.Sleep(time.Second)
		}
	}()
}

func (n ManipulatorNode) transitResource() {
	neededItem, err := n.sender.AskForNeedItem(n.nextNode)
	if err != nil {
		sugar.Sugar.Infof("Error while asking %v for needed item %v", n.nextNode, err)
		return
	}
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
