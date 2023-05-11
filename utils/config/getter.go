package config

import (
	"path/filepath"

	pgx "github.com/jackc/pgx/v5"
)

func GetCommonConfig() *CommonConfig {
	return &config.CommonConfig
}

func GetServerConfig() *ServerConfig {
	return &config.ServerConfig
}

func GetAuthConfig() *AuthConfig {
	return &config.AuthConfig
}

func getDatabaseConfig() *DBConfig {
	return &config.DatabaseConfig
}

func GetConnConfig() (*pgx.ConnConfig, error) {
	dbconfig := getDatabaseConfig()

	connConfig, err := pgx.ParseConfig("")
	if err != nil {
		return nil, err
	}

	connConfig.Host = dbconfig.Host
	connConfig.Port = uint16(dbconfig.Port)
	connConfig.User = dbconfig.Username
	connConfig.Password = dbconfig.Password
	connConfig.Database = dbconfig.Database

	return connConfig, nil
}

func GetS3Config() *S3Config {
	return &config.S3Config
}

// TODO - move it
func GetAbsPrivatePath() (string, error) {
	authConfig := GetAuthConfig()

	absolutePrivatePath, err := filepath.Abs(*&authConfig.PrivatePath)
	if err != nil {
		return "", err
	}

	return absolutePrivatePath, nil
}

func GetAbsPublicPath() (string, error) {
	authConfig := GetAuthConfig()

	absolutePublicPath, err := filepath.Abs(*&authConfig.PublicPath)
	if err != nil {
		return "", err
	}

	return absolutePublicPath, nil
}
