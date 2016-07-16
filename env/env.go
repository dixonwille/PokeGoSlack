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
	//KeyForm is the key used to get the parsed reqest form.
	KeyForm Key = iota
	//KeyArgs is the key for the arguments for the command.
	KeyArgs
	//KeyHelpCmd is the key for which command to display the help for
	KeyHelpCmd
	//KeyCmd is the command that got called
	KeyCmd
	//KeyDB is the instance of the database
	KeyDB
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
