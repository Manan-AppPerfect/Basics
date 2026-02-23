package structs


// struct
type Car struct {
	Model string
	Engine int
	RevModule bool
	Driver1 Driver
}

// embedded struct

type Driver struct {
	Name string
	WDC int
}

// method
func (c Car) Honk() string {
	if c.RevModule {
		return "Grrr start"
	} 
	return "Normal start"
}

func (c *Car) UpdateEngine() {
	c.Engine++
}