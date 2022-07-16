package workers

import (
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type Mine struct {
	production basic.ItemType
	storage    basic.Storage
	basic.BaseWorkerNode
}

var _ basic.WorkerNode = Mine{}

func NewMine(
	nextNode basic.Node, production basic.ItemType, sender basic.Sender,
) Mine {
	res := Mine{
		production: production,
		storage:    basic.NewStorage(),
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		nextNode,
	)
	return res
}

func (m Mine) StartWorker() {
	go m.SendItems()
	go m.produce()
}

func (m Mine) produce() {
	for {
		sugar.Sugar.Infof("Do work %v", m.storage)

		err := m.Storage.Add(m.production)
		sugar.Sugar.Infof("After work. Err %v LocalStore %s", err, m.Storage)

		time.Sleep(time.Second)
	}
}
