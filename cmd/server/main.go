package main

import (
	"fmt"
	"os"

	"github.com/mcctrix/ctrix-social-go-backend/internal/app"
	"github.com/mcctrix/ctrix-social-go-backend/internal/config"
	db "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
	"github.com/mcctrix/ctrix-social-go-backend/internal/server"
)

func main() {

	cfg := config.LoadConfig()

	db, err := db.DBConnect(cfg.DBConfig)
	if err != nil {
		fmt.Println("Error connecting to db: ", err)
		os.Exit(1)
	}
	CheckArgs()

	services := app.BuildApplicationDependencies(db)

	mainRouter := server.NewServer(services)

	err = mainRouter.Listen(":" + cfg.Port)
	if err != nil {
		fmt.Println(err)
	}
}

func CheckArgs() {
	if len(os.Args) == 0 {
		return
	}
	if utils.ContainsString(os.Args, "reset") {
		db.ResetDB()
	}
	if utils.ContainsString(os.Args, "init-db") {
		db.CreateInitialDBStructure()
	}
	if utils.ContainsString(os.Args, "populate-db") {
		db.PopulateDB()
		os.Exit(0)
	}
}
