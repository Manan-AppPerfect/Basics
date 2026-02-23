package arrays

import "fmt"

func Array () {
	// var arr [5]int = [5]int{3,5,6,3,2}
	
	arr := [5]int{0,43,32,4321,24}

	// fmt.Println(arr[0]) first element
	// arr[1] = 20  modify element
	for _, num := range arr {
		fmt.Println(num)
	}
}