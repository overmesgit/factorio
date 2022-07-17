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
	IronGear  ItemType = "IR_GE"

	NoItem  ItemType = "NoItem"
	AnyItem ItemType = "AnyItem"
)
