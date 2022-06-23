package conf 

import (
	"log"
	"strconv"
	"os"
)

const (
	dbPortKey = "DB_PORT"
	dbHostKey = "DB_HOST"
	dbNameKey = "DB_NAME"
	dbUserKey = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	hostKey = "HOST"
	portKey = "PORT"
)

type Config struct {
	DbHost string
	DbPort string 
	DbName string
	DbUser string 
	DbPassword string 
	Host string 
	Port string 
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
 
	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
	}
}

func logAndPanic(envVar string) {
	log.Println("ENV variable not set or value not valid: ", envVar)
	panic(envVar)
}
