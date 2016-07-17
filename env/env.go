package env

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	osPort         string
	osToken        string
	osDBConnString string
	//Port is the port to listen on
	Port string
	//Logger is the main stdout for all logs
	Logger *log.Logger
	//Token the application has access too
	Token string
	//DBConnString is the connection string to the database
	DBConnString string
	//ClientID is the Id of the application
	ClientID string
	//ClientSecret is the secret to identify application with
	ClientSecret string
)

//Key are used for context throughout application
type Key int

const (
	defaultPort = "80"
)
const (
	//KeyReq is used for general context all methods should have
	KeyReq Key = iota
	//KeyCode is used for Auth
	KeyCode
)

func init() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
	godotenv.Load()
	osPort = os.Getenv("PORT")
	if osPort == "" {
		osPort = defaultPort
	}
	osToken = os.Getenv("SLACK_TOKEN")
	osDBConnString = os.Getenv("DATABASE_URL")
	ClientID = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
	flag.StringVar(&Port, "p", osPort, "Port to listen on")
	flag.StringVar(&Token, "t", osToken, "Token that is accepted by applications")
	flag.StringVar(&DBConnString, "d", osDBConnString, "Database connection string")
	flag.Parse()
}
