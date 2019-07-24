package nfl

type (
	GSIS string
	ESB  string
)

func (g GSIS) String() string {
	return string(g)
}
