package nodemap

type Type string
type Direction string

const (
	Up    Direction = "A"
	Down  Direction = "V"
	Left  Direction = "<"
	Right Direction = ">"
)
const (
	IronMine    Type = "MI"
	CoalMine    Type = "MC"
	Belt        Type = "BE"
	Furnace     Type = "FU"
	Manipulator Type = "MA"
)

type ItemType string

const (
	Iron      ItemType = "IRON"
	Coal      ItemType = "COAL"
	IronPlate ItemType = "IRONPLATE"
)
