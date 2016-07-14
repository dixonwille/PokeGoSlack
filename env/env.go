package env

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	osPort string
	//Port is the port to listen on
	Port string
	//Logger is the main stdout for all logs
	Logger *log.Logger
)

const (
	defaultPort = "80"
)

func init() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
	godotenv.Load()
	osPort = os.Getenv("POKE_PORT")
	if osPort == "" {
		osPort = defaultPort
	}
	flag.StringVar(&Port, "p", osPort, "Port to listen on")
	flag.Parse()
}
