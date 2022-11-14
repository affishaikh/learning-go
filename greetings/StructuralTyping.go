package greetings

type Human interface {
	Name() string
}

type Man struct {
	name string
	age int
}

type Woman struct {
	name string
	age int
}

type Zombie struct {
	name string
	age int
}

func (m Man) Name() string {
	return m.name
}

func (w Woman) Name() string {
	return w.name
}

func printHuman(h Human) {
	println(h.Name())
}

func main() {
	printHuman(Man{})
	printHuman(Woman{})
}
