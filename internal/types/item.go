package types

type ItemType string

const (
	general     ItemType = "general"
	weapon      ItemType = "weapon"
	ammunition  ItemType = "ammunition"
	armor       ItemType = "armor"
	conjuration ItemType = "conjuration"
	potion      ItemType = "potion"
)

type Item struct {
	id          string
	name        string
	price       int
	weight      float64
	description string
	type_       ItemType
}
