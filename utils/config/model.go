package config

type ServerConfig struct {
	Port string `json:"port"`
	Host string `json:"host"`
}

type AuthConfig struct {
	PrivatePath string `json:"privatePath"`
	PublicPath  string `json:"publicPath"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	ServerConfig   ServerConfig `json:"serverConfig"`
	AuthConfig     AuthConfig   `json:"authConfig"`
	DatabaseConfig DBConfig     `json:"databaseConfig"`
}
