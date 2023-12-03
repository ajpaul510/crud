package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppSettings struct {
	connection_string string
	port              int
}

func (as *AppSettings) Init(env string) (string, string) {
	read_env(env)

	if cs := os.Getenv("connection_string"); cs == "" {
		log.Fatal("Empty connection string environment variable")
	}
	if p := os.Getenv("port"); p == "" {
		log.Fatal("Empty port environment variable")
	}

	return os.Getenv("connection_string"), os.Getenv("port")
}

func (as *AppSettings) GetConnectionString() string {
	return as.connection_string
}

func (as *AppSettings) GetPort() string {
	return as.connection_string
}

func (as *AppSettings) SetConnectionString(cs string) {
	as.connection_string = cs

}

func (as *AppSettings) SetPort(p int) {
	as.port = p
}

func read_env(env string) {
	s := fmt.Sprintf("config/%s.env", env)
	err := godotenv.Load(s)
	if err != nil {
		s := fmt.Sprintf("Error loading %s", s)
		log.Fatal(s)
	}
}
