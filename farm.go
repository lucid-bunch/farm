// Package farm provides basic examples for how to document Go modules.
package farm

// Farm is an area of land and its buildings, used for growing crops
// and rearing animals.
type Farm struct {
	Sty    *Sty
	Stable *Stable
}

// String returns the string representation of a Farm.
func (f *Farm) String() string {
	return ""
}

// New constructs a Farm with a Stable and a Sty.
func New() *Farm {
	return &Farm{
		Stable: &Stable{},
		Sty:    &Sty{},
	}
}

// Stable is a building in which livestock, especially horses, are kept.
type Stable struct{}

// Horse returns what a horse does.
func (s *Stable) Horse() string {
	return "runs around and says neigh..."
}

// Sty is a small-scale outdoor enclosure for raising domestic pigs as livestock.
type Sty struct{}

// SpiderPig returns what a spiderpig does.
func (s *Sty) SpiderPig() string {
	return "does whatever a spiderpig does..."
}
