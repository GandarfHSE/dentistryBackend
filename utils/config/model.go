package config

type CommonConfig struct {
	Keyword string `json:"keyword"`
}

type HTTPSConfig struct {
	CertFile string `json:"certFile"`
	KeyFile  string `json:"keyFile"`
}

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

type S3Config struct {
	Endpoint  string `json:"endpoint"`
	Bucket    string `json:"bucket"`
	Region    string `json:"region"`
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type Config struct {
	CommonConfig   CommonConfig `json:"commonConfig"`
	HTTPSConfig    HTTPSConfig  `json:"httpsConfig"`
	ServerConfig   ServerConfig `json:"serverConfig"`
	AuthConfig     AuthConfig   `json:"authConfig"`
	DatabaseConfig DBConfig     `json:"databaseConfig"`
	S3Config       S3Config     `json:"s3Config"`
}
