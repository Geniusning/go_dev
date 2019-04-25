package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct{}

func (p *RandomBalance) DoBalance(ints []*Instances) (inst *Instances, err error) {
	if len(ints) == 0 {
		err = errors.New("no instance")
		return
	}

	lens := len(ints)
	index := rand.Intn(lens)
	inst = ints[index]

	return
}
