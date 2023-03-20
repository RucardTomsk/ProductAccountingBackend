package enum

type TypeWeight int

const (
	KG TypeWeight = iota
	G
)

func (i TypeWeight) String() string {
	return [...]string{"kg", "g"}[i]
}

func ParseTypeWeight(typeWightString string) TypeWeight {
	switch typeWightString {
	case KG.String():
		return KG
	case G.String():
		return G
	default:
		return G
	}
}
