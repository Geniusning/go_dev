package balance

import "strconv"

//Instances is shili
type Instances struct {
	host string
	port int
}

//NewInstace is 构造函数
func NewInstace(host string, port int) *Instances {
	return &Instances{
		host: host,
		port: port,
	}
}

func (p *Instances) GetHost() string {
	return p.host
}

func (p *Instances) GetPort() int {
	return p.port
}

func (p *Instances) String() string {
	return p.host + ":" + strconv.Itoa(p.port)
}
