package basic

type ItemCounter struct {
	Type  string
	Count int64
}

type WorkerNode interface {
	ReceiveResource(itemType Item) error
	GetNeededResource() (ItemType, error)
	GetResourceForSend(ItemType) (Item, error)
	StartWorker()
	GetItemCount() []ItemCounter
}

type Sender interface {
	SendItem(adjNode Node, forSend Item) error
	AskForItem(prevNode Node, itemType ItemType) (Item, error)
	AskForNeedItem(nextNode Node) (ItemType, error)
}

type Node struct {
	Row, Col  int32
	NodeType  Type
	Direction Direction
	Hostname  string
}

func NewNode(
	row, col int32, nodeType Type, direction Direction, hostname string,
) Node {

	return Node{Row: row, Col: col, NodeType: nodeType,
		Direction: direction, Hostname: hostname}
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
