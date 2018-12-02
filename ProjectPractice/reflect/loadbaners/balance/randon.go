package balance

import (
"errors"
"math/rand"
)
func init(){

}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	lens := len(insts)
	if lens==0{
		err = errors.New("no backend found")
		return
	}
	index := rand.Intn(lens)
	inst = insts[index]
	return
}