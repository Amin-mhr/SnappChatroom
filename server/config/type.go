package config

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	HttpPort uint `json:"httpPort"`
}
