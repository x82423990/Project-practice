package balance

import "strconv"

type Instance struct {
	host string
	port int
}

func Newinstance(host string, port int) *Instance{
	return &Instance{
		port:port,
		host:host,
	}
}

func (p *Instance) GetHost() string{
	return p.host
}
func (p *Instance) GetPort()int {
	return p.port
}

func (p * Instance)String()string  {
	return p.host +":"+ strconv.Itoa(p.port)
}
