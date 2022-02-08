package animals

type Eagle struct {
	Name string
}

func (c *Eagle) Move() string {
	return "I'm an Eagle named '" + c.Name + "' and I fly"
}
