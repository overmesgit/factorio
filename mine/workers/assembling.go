package workers

import (
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type AssemblingMachine struct {
	basic.BaseWorkerNode
	storage    basic.Storage
	production basic.ItemType
}

var _ basic.WorkerNode = AssemblingMachine{}

func NewAssemblingMachine(
	nextNode basic.Node, production basic.ItemType, sender basic.Sender,
) AssemblingMachine {
	res := AssemblingMachine{
		storage:    basic.NewStorage(),
		production: production,
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		nextNode,
	)
	return res
}

func (n AssemblingMachine) StartWorker() {
	go n.assemble()
}

func (n AssemblingMachine) assemble() {
	for {
		time.Sleep(time.Second)
		storage := n.storage
		if storage.GetCount(basic.IronPlate) == 0 {
			continue
		}

		if storage.IsFull(n.production) {
			continue
		}

		// TODO: errors
		storage.Get(basic.IronPlate)
		storage.Add(n.production)
	}
}

func (n AssemblingMachine) GetResourceForSend(basic.ItemType) (basic.ItemType, error) {
	item, err := n.Storage.Get(n.production)
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", n.Storage.GetItemCount())
		return "", err
	}
	return item, nil
}

func (n AssemblingMachine) GetNeededResource() (basic.ItemType, error) {
	return basic.IronPlate, nil
}
