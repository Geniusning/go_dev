package balance

//
type Balancer interface {
	DoBalance([]*Instances) (*Instances, error)
}
