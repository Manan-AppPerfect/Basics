package conditionals

func Switch(a int) string {

	// switch a {

	// case 1:
	// 	return "Monday"
	// case 2:
	// 	return "Tuesday"
	// case 3:
	// 	return "Wednesday"
	// case 4:
	// 	return "Thursday"
	// case 5:
	// 	return "Friday"
	// default:
	// 	return "Holiday yay"
	// }

	// we can also write like this

	switch {
	case a == 1, a == 2, a == 3, a == 4, a == 5 :
		return "Working day"
	default:
		return "Holiday off day"
	}



	// we can use fallthrough to check the conditions below if our condition is met
}