package config

type Config struct {
	Server ServerConfig `json:"server"`
	Nats   NATSConfig   `json:"nats"`
	DB     DBConfig     `json:"db"`
}

type ServerConfig struct {
	HttpPort uint `json:"httpPort"`
}

type NATSConfig struct {
	HostPort string `json:"hostPort"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}
