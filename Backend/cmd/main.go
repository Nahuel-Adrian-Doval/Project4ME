package main

import (
	"fmt"
	"net/http"

	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/config"
	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/internal/api/router"
	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/internal/pkg"
	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/models"
	"github.com/Nahuel-Adrian-Doval/Project4ME/Backend/utils"
	"github.com/joho/godotenv"
)

func main() {
	_init_()
}

func _init_() {

	loadEnvTask := make(chan bool)
	go loadEnv(loadEnvTask)

	<-loadEnvTask

	dbConf := config.LoadDBConfig()
	serverConf := config.LoadServerConfig()

	startDatabaseTask := make(chan bool)
	go startDatabase(dbConf, startDatabaseTask)

	// Wait to the startDatabase finished
	<-startDatabaseTask

	startServer(serverConf) // Start server
}

func loadEnv(loadEnvTask chan bool) {
	err := godotenv.Load(".env")
	if err != nil {
		utils.SaveLog("Failed to load .env file")
		panic("Failed to load .env file")
	}

	fmt.Println("Loaded .env file!!")
	loadEnvTask <- true
}

func startServer(serverConf models.ServerConfig) {
	mux := http.NewServeMux()

	router.Routes(serverConf.SERVER_PREFIX, mux)

	err := http.ListenAndServe(":"+serverConf.SERVER_PORT, mux)
	if err != nil {
		utils.SaveLog(fmt.Sprintf("Failed to start server: %s", err.Error()))
		panic("Failed to start server")
	}
}

func startDatabase(dbConf models.DBConfig, startDatabaseTask chan bool) {
	_, err := pkg.DBConnection(dbConf)
	if err != nil {
		utils.SaveLog(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
		panic("Failed to connect to database")
	}

	fmt.Println("Connected to database!!")

	startDatabaseTask <- true
}
