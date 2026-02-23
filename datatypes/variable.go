package datatypes

import (
	"fmt"
	"math"
	"strconv"
)

func Variable(){

	/*
	Data Types in go are as follows 

	int int8 int16 int32 int64 for negative(1) & positive(0)
	uint uint8 unint16 uint32 uint64 for only positive value
	float32 float64 float values
	rune: alias for int32
	byte: alias for int8

	1byte - 8bits

	string 
	bool

	use fmt package for formatting
	use strconv package for conversion of string to integers
	use math package for various math resources 
	*/

	x := "232342"
	y, err:= strconv.Atoi(x)

	a := float64(4)
	b := math.Pow(a, 3)

	fmt.Println(y, err, b)

}