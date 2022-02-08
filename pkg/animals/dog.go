package animals

type Dog struct {
	Name string
}

func (d *Dog) Move() string {
	return "I'm a Dog named '" + d.Name + "' and I run"
}
