package generic

import "fmt"

func First[T any](arr []T) T {
	return arr[0]
}

// type Box[T any] struct{
// 	value T
// }


func Generic () {
	fmt.Println(First([]int{1,2,3}))
	fmt.Println(First([]float32{3.32,32.3,3.32,32}))
	fmt.Println(First([]string{"Hi", "This", "manan"}))
}