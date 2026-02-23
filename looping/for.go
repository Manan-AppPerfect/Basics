package looping

import "fmt"

func ForLoop () {
	for idx := 0; idx < 10; idx++ {
		fmt.Println(idx)
	}

	str := "Hello World"

	//you can directly iterate through string in go use Range

	for _, char := range str {
		fmt.Printf("%c", char)
	}

	// _ means ignore this variable

	// can use break and continue 
}