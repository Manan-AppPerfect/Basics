package main

import (
	// "basics/conditionals"
	// "basics/datatypes"
	// "basics/looping"
	// "basics/arrays"
	// "basics/interfaces"
	"basics/generic"
	"basics/pointers"
	"basics/structs"
	"fmt"
)

func main(){
	RB := structs.Car{
		Model: "RB19",
		Engine: 12,
		RevModule: true,
		Driver1: structs.Driver{
			Name: "MV33",
			WDC: 4,
		},
	}

	RB.UpdateEngine()
	fmt.Println(RB)
	// fmt.Println(RB.Driver1.WDC)
	// datatypes.Variable()
	// fmt.Println(conditionals.Compare())
	// fmt.Println(conditionals.IfElse(5,5))
	// fmt.Println(conditionals.Switch(3))
	// looping.ForLoop()
	// looping.WhileLoop()
	// interfaces.InterfaceShapes()
	// arrays.Array()
	pointers.Pointers()
	generic.Generic()

}