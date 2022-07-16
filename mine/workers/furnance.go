package workers

import (
	"github.com/overmesgit/factorio/mine"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type FurnaceNode struct {
	storage basic.Storage
	basic.WorkerNode
}

var _ WorkerNode = FurnaceNode{}

func NewFurnaceNode(
	nextNode Node,
) FurnaceNode {
	res := FurnaceNode{
		storage: basic.NewStorage(),
	}
	sender := basic.NewSender(
		&res.storage,
		mine.NewSender(),
		nextNode,
	)
	res.WorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		IronPlate,
	)
	return res
}

func (n FurnaceNode) StartWorker() {
	go n.SendItems()
	go n.melt()
}

func (n FurnaceNode) melt() {
	for {
		time.Sleep(time.Second)
		storage := n.storage
		if storage.GetCount(Iron) == 0 || storage.GetCount(Coal) == 0 {
			continue
		}

		if storage.IsFull(IronPlate) {
			continue
		}

		// TODO: errors
		storage.Get(Iron)
		storage.Get(Coal)
		storage.Add(IronPlate)
	}
}

func (n FurnaceNode) GetNeededResource() (ItemType, error) {
	if n.storage.GetCount(Iron) > n.storage.GetCount(Coal) {
		return Coal, nil
	} else {
		return Iron, nil
	}
}
