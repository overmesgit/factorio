package workers

import (
	"github.com/overmesgit/factorio/mine"
	"time"
)

type Mine struct {
	Node
	Storage
	mine.Sender
}

func NewMine(
	row, col int32, nodeType Type, direction Direction, production ItemType,
) *Mine {
	return &Mine{Node: NewNode(row, col, nodeType, direction, production)}
}

func (m *Mine) DoWork() {
	go m.produce()
	go m.sendItems()
}

func (m *Mine) produce() {
	for {
		Sugar.Infof("Do work %v", m)

		err := m.Add(m.production)
		Sugar.Infof("After work. Err %v LocalStore %s", err, m.Storage)

		time.Sleep(time.Second)
	}
}

func (m *Mine) sendItems() {
	for {
		time.Sleep(200 * time.Millisecond)
		adjNode := m.getNextNode()
		if adjNode != nil {
			err := m.SendItemFromStore(adjNode)
			if err != nil {
				Sugar.Errorf("err: %v", err)
			}
		}
	}
}
