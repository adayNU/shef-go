package shef

type Receiver struct {
	ip   string
	port string
}

func (r *Receiver) Address() string {
	return r.ip + ":" + r.port
}

func NewReceiver(ip, port string) *Receiver {
	return &Receiver{
		ip:   ip,
		port: port,
	}
}
