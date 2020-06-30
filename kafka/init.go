package kafka

type Kafka struct {
	Ip string
	Port string
	Connect string
}

func NewKafka(ip string, port string, connect string) *Kafka {
	return &Kafka{Ip: ip, Port: port, Connect: connect}
}
