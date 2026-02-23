package looping

import "fmt"

/*
inside go we don't have while loop we can write while loop using for
*/

func WhileLoop(){
	a := 1;
	for a < 10 {
		fmt.Println("Loop")
		a++
	}
}