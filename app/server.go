package app

import (
	"P2/database/seeders"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.InitializeDB(dbConfig)
	server.InitializeRoute()
	seeders.DBSeed(server.DB)
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) InitializeDB(dbConfig DBConfig) {
	var err error

	if dbConfig.DBDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBHost,
			dbConfig.DBPort,
			dbConfig.DBName,
		)
		server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbConfig.DBHost,
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBName,
			dbConfig.DBPort,
		)
		server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	for _, model := range RegistryModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Database Migration Successfully")
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatal("Error on Load .env File")
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppEnv = os.Getenv("APP_ENV")
	appConfig.AppPort = os.Getenv("APP_PORT")

	dbConfig.DBHost = os.Getenv("DB_HOST")
	dbConfig.DBUser = os.Getenv("DB_USER")
	dbConfig.DBPassword = os.Getenv("DB_PASSWORD")
	dbConfig.DBName = os.Getenv("DB_NAME")
	dbConfig.DBPort = os.Getenv("DB_PORT")
	dbConfig.DBDriver = os.Getenv("DB_DRIVER")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
