package config

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

type Config struct {
	Servername string
	Port       int
	Env        string
	Cors       []string
	Roles      struct {
		Default    string
		Additional []string
	}
	BaseRouting string
	DB          struct {
		DSN  string
		Name string
	}
	JWT    []byte
	Cookie struct {
		Name     string
		SameSite string
	}
}

type userContextKey int

type AppConfig struct {
	Version         string
	ServerStartTime time.Time
	Config          Config
	Logger          *log.Logger
	Validator       *validator.Validate
	UserContextKey  userContextKey
}

func Init() AppConfig {
	var app AppConfig

	flag.StringVar(&app.Config.Env, "env", "dev", "the app environment")

	flag.StringVar(&app.Config.Servername, "apiservername", "http://localhost", "the api server name")
	flag.IntVar(&app.Config.Port, "port", 8080, "the port")

	flag.StringVar(&app.Config.DB.DSN, "dsn", "mongodb://localhost:27017", "the db dsn")
	flag.StringVar(&app.Config.DB.Name, "dbName", "basic-auth", "the name of the used database")

	var jwt string
	flag.StringVar(&jwt, "jwt", "wonderfulsecretphrase", "the jwt token secret")

	var cors string
	flag.StringVar(&cors, "cors", "http://* https://*", "the by cors allowed origins")

	flag.StringVar(&app.Config.Roles.Default, "defaultRole", "guest", "the default role when creating a new user")
	var addRoles string
	flag.StringVar(&addRoles, "addRoles", "guest", "additional roles (admin always exists)")

	flag.StringVar(&app.Config.BaseRouting, "baseRouting", "users", "the base routing for the users endpoints")

	flag.StringVar(&app.Config.Cookie.Name, "cookieName", "basic-auth", "the name of the cookie")
	flag.StringVar(&app.Config.Cookie.SameSite, "cookieSameSite", "lax", "the cookie same site policy")

	flag.Parse()
	app.Config.JWT = []byte(jwt)
	app.Config.Cors = strings.Split(cors, " ")
	app.Config.Roles.Additional = strings.Split(addRoles, " ")

	app.ServerStartTime = time.Now()
	flag.StringVar(&app.Version, "version", "1.0.0", "the api version")

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app.Logger = logger

	app.Validator = validator.New()

	app.UserContextKey = 0

	return app
}
