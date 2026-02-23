package pointers

import "fmt"

func Pointers() {
	a := 20

	p := &a

	fmt.Println(*p)
}

/*
func change(x *int) {
    *x = 20
}

func main() {
    x := 10
    change(&x)
    fmt.Println(x) // now 20
}
*/