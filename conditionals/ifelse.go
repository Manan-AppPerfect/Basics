package conditionals

func IfElse(x , y int) string {

	if x < y {
		return "y>x"
	} else if x == y {
		return "equal"
	}
	return "x>y"
}