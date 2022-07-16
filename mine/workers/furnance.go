package workers

import (
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type FurnaceNode struct {
	storage basic.Storage
	basic.BaseWorkerNode
}

var _ basic.WorkerNode = FurnaceNode{}

func NewFurnaceNode(
	nextNode basic.Node, sender basic.Sender,
) FurnaceNode {
	res := FurnaceNode{
		storage: basic.NewStorage(),
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		nextNode,
		basic.IronPlate,
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
		if storage.GetCount(basic.Iron) == 0 || storage.GetCount(basic.Coal) == 0 {
			continue
		}

		if storage.IsFull(basic.IronPlate) {
			continue
		}

		// TODO: errors
		storage.Get(basic.Iron)
		storage.Get(basic.Coal)
		storage.Add(basic.IronPlate)
	}
}

func (n FurnaceNode) GetNeededResource() (basic.ItemType, error) {
	if n.storage.GetCount(basic.Iron) > n.storage.GetCount(basic.Coal) {
		return basic.Coal, nil
	} else {
		return basic.Iron, nil
	}
}
