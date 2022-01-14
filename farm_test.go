package farm_test

import (
	"fmt"

	"github.com/lucid-bunch/farm"
)

func ExampleNew() {
	f := farm.New()

	fmt.Println(f)
	// Output:
}

func ExampleStable_Horse() {
	f := farm.New()
	h := f.Stable.Horse()

	fmt.Println(h)
	// Output:
	// runs around and says neigh...
}

func ExampleSty_SpiderPig() {
	f := farm.New()
	s := f.Sty.SpiderPig()

	fmt.Println(s)
	// Output:
	// does whatever a spiderpig does...
}
