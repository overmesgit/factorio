package workers

import (
	"github.com/overmesgit/factorio/mine"
	"github.com/overmesgit/factorio/mine/sugar"
	"github.com/overmesgit/factorio/mine/workers/basic"
	"time"
)

type Mine struct {
	production ItemType
	storage    basic.Storage
	basic.WorkerNode
}

var _ WorkerNode = Mine{}

func NewMine(
	nextNode Node, production ItemType,
) Mine {
	res := Mine{
		production: production,
		storage:    basic.NewStorage(),
	}
	sender := basic.NewSender(
		&res.storage,
		mine.NewSender(),
		nextNode,
	)
	res.WorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		production,
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
