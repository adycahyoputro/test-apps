package server

import (
	"backend/config"
	"backend/router"
	"log"
)

func Run() error {
	db, err := config.InitDB()
	if err != nil {
		return err
	}
	defer config.CloseDB(db)

	router := router.SetupRouter(db)

	// serverAddress := utils.GetEnv("SERVER_ADDRESS")
	serverAddress := "192.168.157.160:8082"
	log.Printf("server is running on address : %s\n", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		return err
	}
	return nil
}