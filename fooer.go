package testgover

import "fmt"

const (
	foo = "foo"
	bar = "bar"
	baz = "baz"
)

func Fooer() {
	fmt.Println(foo)
}

func Barer() {
	fmt.Println(bar)
}

func Bazer() {
	fmt.Println(baz)
}
