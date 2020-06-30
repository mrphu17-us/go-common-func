package database

type Database struct {
	Ip string
	Port string
	Connect string
}

func NewDatabase(ip string, port string, connect string) *Database {
	return &Database{Ip: ip, Port: port, Connect: connect}
}
