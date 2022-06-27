package conf 

import (
	"strconv"
	"os"

	"github.com/rs/zerolog/log"
)

const (
	dbPortKey = "DB_PORT"
	dbHostKey = "DB_HOST"
	dbNameKey = "DB_NAME"
	dbUserKey = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	hostKey = "PIC2TEXT_HOST"
	portKey = "PIC2TEXT_PORT"
	jwtSecretKey = "JWT_SECRET"
)

type Config struct {
	DbHost string
	DbPort string 
	DbName string
	DbUser string 
	DbPassword string 
	Host string 
	Port string 
	JwtSecret string 
}

func NewConfig() Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logAndPanic(hostKey)
	}

	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		if _, err := strconv.Atoi(port); err != nil {
			logAndPanic(portKey)
		}
	}

	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logAndPanic(dbHostKey)
	}
	
	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		if _, err := strconv.Atoi(dbPort); err != nil {
		logAndPanic(dbPortKey)
		}
	}
	
	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logAndPanic(dbNameKey)
	}
	
	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logAndPanic(dbUserKey)
	}
	
	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logAndPanic(dbPasswordKey)
	}

	jwtSecret, ok := os.LookupEnv(jwtSecretKey)
	if !ok || jwtSecret == "" {
		logAndPanic(jwtSecretKey)
	}
 
	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
	}
}

func logAndPanic(envVar string) {
	log.Panic().Str("envVar", envVar).Msg("ENV variable not set or value not valid")
}
