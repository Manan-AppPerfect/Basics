package pointers

import "fmt"

func Pointers() {
	// a := 20

	// p := &a

	// fmt.Println(*p)

    paul := User{
        Name: "Paul",
    }
    Change(&paul)
    fmt.Println(paul)
}


type User struct {
    Name string
}

func Change(u *User){
    u.Name = "Saul"
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