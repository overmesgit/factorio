package basic

import (
	"errors"
	"github.com/overmesgit/factorio/mine/sugar"
	"strings"
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

func LogResource(item Item, offset int) {
	offsetStr := strings.Repeat("     ", offset)
	sugar.Sugar.Infof("%s^ Item %v %v", offsetStr, item.Id, item.ItemType)
	for _, ingredient := range item.Ingredients {
		LogResource(*ingredient, offset+1)
	}
}

func (d BaseWorkerNode) ReceiveResource(item Item) error {
	LogResource(item, 0)
	return d.Storage.Add(item)
}

func (d BaseWorkerNode) GetNeededResource() (ItemType, error) {
	return "", errors.New("nothing needed")
}

func (d BaseWorkerNode) GetResourceForSend(itemType ItemType) (Item, error) {
	var forSend Item
	var err error
	if itemType == AnyItem {
		forSend, err = d.Storage.GetAnyItem()
	} else {
		forSend, err = d.Storage.Get(itemType)
	}
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", d.Storage.GetItemCount())
		return Item{ItemType: NoItem}, err
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
