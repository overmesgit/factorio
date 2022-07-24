package workers

import (
	"fmt"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type AssemblingMachine struct {
	basic.BaseWorkerNode
	storage    basic.Storage
	production basic.ItemType
	myNode     basic.Node
}

var _ basic.WorkerNode = AssemblingMachine{}

func NewAssemblingMachine(
	myNode basic.Node, production basic.ItemType, sender basic.Sender,
) AssemblingMachine {
	res := AssemblingMachine{
		storage:    basic.NewStorage(),
		production: production,
		myNode:     myNode,
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		myNode.GetNextNode(),
	)
	return res
}

func (n AssemblingMachine) StartWorker() {
	go n.assemble()
}

func (n AssemblingMachine) assemble() {
	counter := 0
	for {
		time.Sleep(time.Second)
		storage := n.storage
		if storage.GetCount(basic.IronPlate) == 0 {
			continue
		}

		if storage.IsFull(n.production) {
			continue
		}

		resource, err := storage.Get(basic.IronPlate)
		if err != nil {
			sugar.Sugar.Errorf("Can't take resource %v %v.", err, n.Storage.GetItemCount())
			continue
		}
		newItem := basic.Item{
			ItemType:    n.production,
			Id:          fmt.Sprintf("%s-%s-%v", n.myNode.Hostname, n.production, counter),
			Parents:     []string{resource.Id},
			Ingredients: []*basic.Item{&resource},
		}
		counter++

		err = storage.Add(newItem)
		if err != nil {
			sugar.Sugar.Errorf("Can't add item %v %v.", err, n.Storage.GetItemCount())
			continue
		}
	}
}

func (n AssemblingMachine) GetResourceForSend(basic.ItemType) (basic.Item, error) {
	item, err := n.Storage.Get(n.production)
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", n.Storage.GetItemCount())
		return basic.Item{ItemType: basic.NoItem}, err
	}
	return item, nil
}

func (n AssemblingMachine) GetNeededResource() (basic.ItemType, error) {
	return basic.IronPlate, nil
}
