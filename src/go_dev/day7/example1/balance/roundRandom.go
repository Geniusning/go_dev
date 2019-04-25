package balance

import (
	"errors"
)

type RoundBalance struct {
	curindex int
}

func (p *RoundBalance) DoBalance(insts []*Instances) (inst *Instances, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(insts)
	inst = insts[p.curindex]
	p.curindex = (p.curindex + 1) % lens

	return
}
