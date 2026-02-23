package conditionals


/*

comparison operators are

>
<
>=
<=
!=
==

logical operators are 

||
&&
!

*/

func Compare() string{
	a := 32
	b := 342

	if a > b {
		return "a is bigger"
	} 
	return "b is bigger"
}