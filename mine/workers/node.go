package workers

import (
	"github.com/overmesgit/factorio/mine"
)

var Sugar = mine.Sugar

type Node struct {
	row, col   int32
	nodeType   Type
	production ItemType
	direction  Direction
}

func NewNode(
	row, col int32, nodeType Type, direction Direction, production ItemType,
) Node {
	return Node{row: row, col: col, nodeType: nodeType,
		direction: direction, production: production}
}

var directionIndex = map[Direction][]int32{
	//  ROW / COL
	"A": {-1, 0},
	"V": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func (n *Node) GetPrevNode() Node {
	nextNode := n.GetNextNode()
	prevRowOff, prevColOff := n.row-nextNode.row, n.col-nextNode.col
	prevRow, prevCol := n.row+prevRowOff, n.col+prevColOff
	return Node{row: prevRow, col: prevCol}
}

func (n *Node) GetNextNode() Node {
	offset, ok := directionIndex[n.direction]
	if !ok {
		offset = []int32{1, 0}
	}
	adjRow, adjCol := n.row+offset[0], n.col+offset[1]
	return Node{row: adjRow, col: adjCol}
}
