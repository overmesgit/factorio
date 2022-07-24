package workers

import (
	"fmt"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type Mine struct {
	myNode     basic.Node
	production basic.ItemType
	storage    basic.Storage
	basic.BaseWorkerNode
}

var _ basic.WorkerNode = Mine{}

func NewMine(
	myNode basic.Node, production basic.ItemType, sender basic.Sender,
) Mine {
	res := Mine{
		myNode:     myNode,
		production: production,
		storage:    basic.NewStorage(),
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		myNode.GetNextNode(),
	)
	return res
}

func (m Mine) StartWorker() {
	go m.SendItems()
	go m.produce()
}

func (m Mine) produce() {
	counter := 1

	for {
		sugar.Sugar.Infof("Do work %v", m.storage)

		newItem := basic.Item{
			ItemType:    m.production,
			Id:          fmt.Sprintf("%s-%s-%v", m.myNode.Hostname, m.production, counter),
			Parents:     nil,
			Ingredients: nil,
		}

		counter++
		err := m.Storage.Add(newItem)
		sugar.Sugar.Infof("After work. Err %v LocalStore %s", err, m.Storage)

		time.Sleep(time.Second)
	}
}
