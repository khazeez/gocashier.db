package server

import (
	"fmt"

	"gocashier.db/api"
	"gocashier.db/config"
	"gocashier.db/pkg"
)

func RunServer() error {
	db, err := config.InitDb()
	if err != nil {
		panic(err)
	}

	defer config.CloseDb(db)

	router := api.Router(db)

	serverAddress := pkg.Load().AppPort
	fmt.Println("Server is running on port: ", serverAddress)

	if err := router.Run(serverAddress); err != nil {
		return err
	}

	return nil
}
