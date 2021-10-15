// Package mondb defines RagnarokTravels Monsters based on JSON-format in https://account.ragnaroktravels.com/mondrops.
package mondb

type Monster struct {
	Name string
	Items
	Aran int `json:"aRan"`
	Info
	DBName []string `json:"db_name"`
}

type Items struct {
	Item1    []string `json:"item1"`
	Percent1 int      `json:"percent1"`
	Item2    []string `json:"item2"`
	Percent2 int      `json:"percent2"`
	Item3    []string `json:"item3"`
	Percent3 int      `json:"percent3"`
	Item4    []string `json:"item4"`
	Percent4 int      `json:"percent4"`
	Item5    []string `json:"item5"`
	Percent5 int      `json:"percent5"`
	Item6    []string `json:"item6"`
	Percent6 int      `json:"percent6"`
	Item7    []string `json:"item7"`
	Percent7 int      `json:"percent7"`
	Item8    []string `json:"item8"`
	Percent8 int      `json:"percent8"`
}

type Info struct {
	Stats
	AttackRange
	Def int `json:"def"`
	Experience
	Inc          int            `json:"inc"`
	As           int            `json:"as"`
	Es           int            `json:"es"`
	Mspeed       int            `json:"Mspeed"`
	RechargeTime int            `json:"rechargeTime"`
	AttackedMT   int            `json:"attackedMT"`
	AttackMT     int            `json:"attackMT"`
	Property     Property       `json:"property"`
	Size         Size           `json:"scale"`
	Class        Classification `json:"class"`
	Race         Race           `json:"race"`
	Mdef         int            `json:"mdef"`
	PetInfo
}

type Stats struct {
	Level int `json:"LV"`
	HP    int
	SP    int
	Str   int `json:"str"`
	Int   int `json:"int"`
	Vit   int `json:"vit"`
	Dex   int `json:"dex"`
	Agi   int `json:"agi"`
	Luk   int `json:"luk"`
}

type AttackRange struct {
	Low  int `json:"atk1"`
	High int `json:"atk2"`
}

type Experience struct {
	Base int `json:"exp"`
	Job  int `json:"jexp"`
}

type PetInfo struct {
	TamingItem string `json:"tamingitem"`
	FoodItem   string `json:"fooditem"`
}

type Property int

func (p Property) String() string {
	switch p {
	case NeutralSpecial, Neutral1:
		return "Neutral 1"
	case Neutral2:
		return "Neutral 2"
	case Neutral3:
		return "Neutral 3"
	case Neutral4:
		return "Neutral 4"
	case Water1:
		return "Water 1"
	case Water2:
		return "Water 2"
	case Water3:
		return "Water 3"
	case Water4:
		return "Water 4"
	case Earth1:
		return "Earth 1"
	case Earth2:
		return "Earth 2"
	case Earth3:
		return "Earth 3"
	case Earth4:
		return "Earth 4"
	case Fire1:
		return "Fire 1"
	case Fire2:
		return "Fire 2"
	case Fire3:
		return "Fire 3"
	case Fire4:
		return "Fire 4"
	case Wind1:
		return "Wind 1"
	case Wind2:
		return "Wind 2"
	case Wind3:
		return "Wind 3"
	case Wind4:
		return "Wind 4"
	case Poison1:
		return "Poison 1"
	case Poison2:
		return "Poison 2"
	case Poison3:
		return "Poison 3"
	case Poison4:
		return "Poison 4"
	case Holy1:
		return "Holy 1"
	case Holy2:
		return "Holy 2"
	case Holy3:
		return "Holy 3"
	case Holy4:
		return "Holy 4"
	case Shadow1:
		return "Shadow 1"
	case Shadow2:
		return "Shadow 2"
	case Shadow3:
		return "Shadow 3"
	case Shadow4:
		return "Shadow 4"
	case Ghost1:
		return "Ghost 1"
	case Ghost2:
		return "Ghost 2"
	case Ghost3:
		return "Ghost 3"
	case Ghost4:
		return "Ghost 4"
	case Undead1:
		return "Undead 1"
	case Undead2:
		return "Undead 2"
	case Undead3:
		return "Undead 3"
	case Undead4:
		return "Undead 4"
	default:
		return ""
	}
}

const (
	// NeutralSpecial is Neutral1 technically, but reserved for special monsters such as WOE Guardians and Treasure Boxes.
	NeutralSpecial Property = 0
	Neutral1                = 20
	Neutral2                = 40
	Neutral3                = 60
	Neutral4                = 80
	Water1                  = 21
	Water2                  = 41
	Water3                  = 61
	Water4                  = 81
	Earth1                  = 22
	Earth2                  = 42
	Earth3                  = 62
	Earth4                  = 82
	Fire1                   = 23
	Fire2                   = 43
	Fire3                   = 63
	Fire4                   = 83
	Wind1                   = 24
	Wind2                   = 44
	Wind3                   = 64
	Wind4                   = 84
	Poison1                 = 25
	Poison2                 = 45
	Poison3                 = 65
	Poison4                 = 85
	Holy1                   = 26
	Holy2                   = 46
	Holy3                   = 66
	Holy4                   = 86
	Shadow1                 = 27
	Shadow2                 = 47
	Shadow3                 = 67
	Shadow4                 = 87
	Ghost1                  = 28
	Ghost2                  = 48
	Ghost3                  = 68
	Ghost4                  = 88
	Undead1                 = 29
	Undead2                 = 49
	Undead3                 = 69
	Undead4                 = 89
)

type Size int

func (s Size) String() string {
	switch s {
	case Small:
		return "Small"
	case Medium:
		return "Medium"
	case Large:
		return "Large"
	default:
		return ""
	}
}

const (
	Small Size = iota
	Medium
	Large
)

type Classification int

const (
	Mob  Classification = iota
	Boss                // Currently the ragnarok travels DB doesn't provide granularity between mini-boss and MVP.
)

type Race int

func (r Race) String() string {
	switch r {
	case Formless:
		return "Formless"
	case Undead:
		return "Undead"
	case Brute:
		return "Brute"
	case Plant:
		return "Plant"
	case Insect:
		return "Insect"
	case Fish:
		return "Fish"
	case Demon:
		return "Demon"
	case DemiHuman:
		return "Demi-Human"
	case Angel:
		return "Angel"
	case Dragon:
		return "Dragon"
	default:
		return ""
	}
}

const (
	Formless Race = iota
	Undead
	Brute
	Plant
	Insect
	Fish
	Demon
	DemiHuman
	Angel
	Dragon
)
