package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var mgr = BalanceMgr{
	allBalancer: make(map[string]Balancer),
}

func (p *BalanceMgr) registerBalance(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalance(name string, b Balancer) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Instances) (inst *Instances, err error) {
	balancer, ok := mgr.allBalancer[name]
	if !ok {
		err = fmt.Errorf("not found %s balancer", name)
		return
	}
	inst, err = balancer.DoBalance(insts)
	return
}
