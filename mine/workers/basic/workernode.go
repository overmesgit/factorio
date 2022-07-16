package basic

import (
	"errors"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers"
	"time"
)

type WorkerNode struct {
	*Storage
	sender     Sender
	production workers.ItemType
}

var _ workers.WorkerNode = WorkerNode{}

func NewWorkerNode(storage *Storage, sender Sender, production workers.ItemType) WorkerNode {
	return WorkerNode{Storage: storage, sender: sender, production: production}
}

func (d WorkerNode) ReceiveResource(itemType workers.ItemType) error {
	return d.Storage.Add(itemType)
}

func (d WorkerNode) GetNeededResource() (workers.ItemType, error) {
	return "", errors.New("nothing needed")
}

func (d WorkerNode) GetResourceForSend() (workers.ItemType, error) {
	item, err := d.Storage.Get(d.production)
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", d.Storage.GetItemCount())
		return "", err
	}
	return item, nil
}

func (d WorkerNode) StartWorker() {
	go d.SendItems()
}

func (d WorkerNode) SendItems() {
	for {
		sugar.Sugar.Infof("Send items from store %v", d.Storage)

		err := d.sender.SendItemFromStore()
		if err != nil {
			sugar.Sugar.Errorf("err: %v", err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
