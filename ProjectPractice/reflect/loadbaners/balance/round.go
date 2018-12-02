package balance

import (
	"errors"
)

type RoundRobinBlance struct {
	curIndex int
}

func (p *RoundRobinBlance) DoBanlance(insts []*Instance, key ...string) (inst *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("no found backen instance")
		return
	}
	if p.curIndex > lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex +1) % lens
	return
}
