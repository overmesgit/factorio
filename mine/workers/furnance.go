package workers

import (
	"fmt"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type FurnaceNode struct {
	storage basic.Storage
	basic.BaseWorkerNode
	myNode basic.Node
}

var _ basic.WorkerNode = FurnaceNode{}

func NewFurnaceNode(
	myNode basic.Node, sender basic.Sender,
) FurnaceNode {
	res := FurnaceNode{
		storage: basic.NewStorage(),
		myNode:  myNode,
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		myNode.GetNextNode(),
	)
	return res
}

func (n FurnaceNode) StartWorker() {
	go n.melt()
}

func (n FurnaceNode) melt() {
	counter := 0
	for {
		time.Sleep(time.Second)
		storage := n.storage
		if storage.GetCount(basic.Iron) == 0 || storage.GetCount(basic.Coal) == 0 {
			continue
		}

		if storage.IsFull(basic.IronPlate) {
			continue
		}

		iron, err := storage.Get(basic.Iron)
		if err != nil {
			sugar.Sugar.Errorf("Can't get item %v %v.", err, n.Storage.GetItemCount())
			continue
		}
		coal, err := storage.Get(basic.Coal)
		if err != nil {
			sugar.Sugar.Errorf("Can't get item %v %v.", err, n.Storage.GetItemCount())
			return
		}
		newIronPlate := basic.Item{
			ItemType:    basic.IronPlate,
			Id:          fmt.Sprintf("%s-%s-%v", n.myNode.Hostname, basic.IronPlate, counter),
			Parents:     []string{iron.Id, coal.Id},
			Ingredients: []*basic.Item{&iron, &coal},
		}
		counter++

		storage.Add(newIronPlate)
	}
}

func (n FurnaceNode) GetResourceForSend(basic.ItemType) (basic.Item, error) {
	item, err := n.Storage.Get(basic.IronPlate)
	if err != nil {
		sugar.Sugar.Infof("Nothing to give %v.", n.Storage.GetItemCount())
		return basic.Item{ItemType: basic.NoItem}, err
	}
	return item, nil
}

func (n FurnaceNode) GetNeededResource() (basic.ItemType, error) {
	if n.storage.GetCount(basic.Iron) > n.storage.GetCount(basic.Coal) {
		return basic.Coal, nil
	} else {
		return basic.Iron, nil
	}
}
