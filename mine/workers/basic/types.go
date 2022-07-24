package basic

type Type string
type Direction string

const (
	Up    Direction = "A"
	Down  Direction = "V"
	Left  Direction = "<"
	Right Direction = ">"
)

const (
	Mine              Type = "MINE"
	Belt              Type = "BELT"
	Furnace           Type = "FURNACE"
	Manipulator       Type = "MANIPULATOR"
	AssemblingMachine Type = "ASSEMBLING_MACHINE"
)

type ItemType string

const (
	Iron      ItemType = "IRON"
	Coal      ItemType = "COAL"
	Copper    ItemType = "COPPER"
	IronPlate ItemType = "IR_PL"
	IronGear  ItemType = "GEAR"

	NoItem  ItemType = "NoItem"
	AnyItem ItemType = "AnyItem"
)

type Item struct {
	ItemType    ItemType `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id          string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Parents     []string `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	Ingredients []*Item  `protobuf:"bytes,4,rep,name=ingredients,proto3" json:"ingredients,omitempty"`
}
