package config

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

type database struct {
	URL string
}

type jwt struct {
	Secret string
	Issuer string
}

type Config struct {
	Database database
	JWT      jwt
}

var (
	dbUser                 = mustGetenv("DB_USER")                  // e.g. 'my-db-user'
	dbPwd                  = mustGetenv("DB_PASS")                  // e.g. 'my-db-password'
	instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	dbName                 = mustGetenv("DB_NAME")                  // e.g. 'my-database'
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func LoadEnv(fileName string) {
	re := regexp.MustCompile(`^(.*` + "fitstackapi" + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/` + fileName)
	if err != nil {
		godotenv.Load()
	}
}

func New() *Config {

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	return &Config{
		Database: database{
			URL: dbURI,
		},
		JWT: jwt{
			Secret: os.Getenv("JWT_SECRET"),
			Issuer: os.Getenv("DOMAIN"),
		},
	}
}
