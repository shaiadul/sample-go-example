package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
}

var configuration *Config

func loadConfig() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Faild to load .env:", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("http port is required")
		os.Exit(1)
	}

	httpPortNum, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("prot must be a number")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("jwt secret is required")
		os.Exit(1)
	}

	configuration = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortNum),
		JWTSecret:   jwtSecret,
	}

}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
