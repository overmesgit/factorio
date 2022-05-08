package workers

import (
	"github.com/overmesgit/factorio/mine"
	"time"
)

type ManipulatorNode struct {
	node    Node
	storage mine.Storage
}

func NewManipulator(
	row, col int32, nodeType Type, direction Direction, production ItemType,
) *ManipulatorNode {
	return &ManipulatorNode{node: NewNode(row, col, nodeType, direction, production)}
}

func (m ManipulatorNode) DoWork() {
	for {
		neededItem, err := s.askForNeedItem(s.getNextNode())
		if err == nil {
			item, err := s.askForItem(s.getPrevNode(), ItemType(Type), false)
			if err == nil {
				err = s.sendItem(s.getNextNode(), item)
			}
		}
		time.Sleep(time.Second)
	}
}
