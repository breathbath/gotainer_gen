package definition

type Conn struct {
	internalState string
}

func (c *Conn) Destroy() error {
	c.internalState = ""
	return nil
}