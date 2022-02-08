package animals

type Orca struct {
	Name string
}

func (o *Orca) Move() string {
	return "I'm an Orca named '" + o.Name + "' and I swim"
}
