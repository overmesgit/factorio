package basic

import (
	"errors"
	"github.com/overmesgit/factorio/mine/sugar"
	"time"
)

type BaseWorkerNode struct {
	*Storage
	sender   Sender
	nextNode Node
}

var _ WorkerNode = BaseWorkerNode{}

func NewWorkerNode(
	storage *Storage, sender Sender, nextNode Node,
) BaseWorkerNode {
	return BaseWorkerNode{Storage: storage, sender: sender, nextNode: nextNode}
}

func (d BaseWorkerNode) ReceiveResource(itemType ItemType) error {
	return d.Storage.Add(itemType)
}

func (d BaseWorkerNode) GetNeededResource() (ItemType, error) {
	return "", errors.New("nothing needed")
}

func (d BaseWorkerNode) GetResourceForSend(item ItemType) (ItemType, error) {
	var forSend ItemType
	var err error
	if item == AnyItem {
		forSend, err = d.Storage.GetAnyItem()
	} else {
		forSend, err = d.Storage.Get(item)
	}
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", d.Storage.GetItemCount())
		return "", err
	}
	return forSend, nil
}

func (d BaseWorkerNode) StartWorker() {
	go d.SendItems()
}

func (d BaseWorkerNode) SendItems() {
	for {
		sugar.Sugar.Infof("Send items from store %v", d.Storage)

		err := d.SendItemFromStore()
		if err != nil {
			sugar.Sugar.Errorf("err: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func (d BaseWorkerNode) SendItemFromStore() error {
	sugar.Sugar.Infof("Send items. Current store: %v", d.Storage)
	forSend, err := d.Storage.GetAnyItem()

	if err != nil {
		sugar.Sugar.Infof("Nothing to send: " + err.Error())
		return err
	}

	err = d.sender.SendItem(
		d.nextNode, forSend,
	)
	if err != nil {
		sugar.Sugar.Warnf("Could send item %v to %v: %v", forSend, d.nextNode, err)
		err := d.Storage.Add(forSend)
		if err != nil {
			sugar.Sugar.Warnf("Could not stack item back %v %v", forSend, err)
		}
		return err
	}

	return nil
}
