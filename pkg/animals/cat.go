package animals

type Cat struct {
	Name string
}

func (c *Cat) Move() string {
	return "I'm a Cat named '" + c.Name + "' and I leap"
}
