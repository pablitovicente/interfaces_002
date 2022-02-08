package animals

type Kangaroo struct {
	Name string
}

func (k *Kangaroo) Move() string {
	return "I'm a Kangaroo named '" + k.Name + "' and I jump"
}
