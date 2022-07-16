package workers

import (
	"github.com/overmesgit/factorio/mine/workers/basic"
)

type Belt struct {
	basic.BaseWorkerNode
	nextNode basic.Node
	sender   basic.Sender
	storage  basic.Storage
}

var _ basic.WorkerNode = Belt{}

func NewBelt(
	nextNode basic.Node, sender basic.Sender,
) Belt {
	res := Belt{
		nextNode: nextNode,
		storage:  basic.NewStorage(),
	}
	res.BaseWorkerNode = basic.NewWorkerNode(
		&res.storage,
		sender,
		nextNode,
	)
	return res
}

func (n Belt) GetNeededResource() (basic.ItemType, error) {
	return basic.AnyItem, nil
}
