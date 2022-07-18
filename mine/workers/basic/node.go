package basic

type ItemCounter struct {
	Type  string
	Count int64
}

type WorkerNode interface {
	ReceiveResource(itemType ItemType) error
	GetNeededResource() (ItemType, error)
	GetResourceForSend(ItemType) (ItemType, error)
	StartWorker()
	GetItemCount() []ItemCounter
}

type Sender interface {
	SendItem(adjNode Node, forSend ItemType) error
	AskForItem(prevNode Node, itemType ItemType) (ItemType, error)
	AskForNeedItem(nextNode Node) (ItemType, error)
}

type Node struct {
	Row, Col  int32
	NodeType  Type
	Direction Direction
}

func NewNode(
	row, col int32, nodeType Type, direction Direction,
) Node {
	return Node{Row: row, Col: col, NodeType: nodeType,
		Direction: direction}
}

var directionIndex = map[Direction][]int32{
	//  ROW / COL
	"^": {-1, 0},
	"V": {1, 0},
	"<": {0, -1},
	">": {0, 1},
}

func (n *Node) GetPrevNode() Node {
	nextNode := n.GetNextNode()
	prevRowOff, prevColOff := n.Row-nextNode.Row, n.Col-nextNode.Col
	prevRow, prevCol := n.Row+prevRowOff, n.Col+prevColOff
	return Node{Row: prevRow, Col: prevCol}
}

func (n *Node) GetNextNode() Node {
	offset, ok := directionIndex[n.Direction]
	if !ok {
		offset = []int32{1, 0}
	}
	adjRow, adjCol := n.Row+offset[0], n.Col+offset[1]
	return Node{Row: adjRow, Col: adjCol}
}
