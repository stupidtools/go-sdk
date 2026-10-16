package target

type Target string

const (
	DEFAULT Target = "DEFAULT"
	NWS     Target = "NWS"
	VAPS    Target = "VAPS"
	Transit Target = "Transit"
	Portico Target = "Portico"
	Realex  Target = "Realex"
	GP_API  Target = "GP_API"
	GNAP    Target = "GNAP"
	NTS     Target = "NTS"
	Genius  Target = "Genius"
)

func (t Target) LongValue() int64 {
	ordinals := map[Target]int{
		DEFAULT: 0,
		NWS:     1,
		VAPS:    2,
		Transit: 3,
		Portico: 4,
		Realex:  5,
		GP_API:  6,
		GNAP:    7,
		NTS:     8,
		Genius:  9,
	}
	if ordinal, ok := ordinals[t]; ok {
		return 1 << ordinal
	}
	return 0
}
