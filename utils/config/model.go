package config

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type AuthConfig struct {
	PrivatePath string `json:"privatePath"`
	PublicPath  string `json:"publicPath"`
}

type Config struct {
	ServerConfig ServerConfig `json:"serverConfig"`
	AuthConfig   AuthConfig   `json:"authConfig"`
}
