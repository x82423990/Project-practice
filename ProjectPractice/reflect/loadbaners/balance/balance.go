package balance


type Balancer interface {
	DoBanlance([]*Instance, ...string)
}
