package main

import (
	"fmt"
	"os"

	"github.com/mcctrix/ctrix-social-go-backend/internal/config"
	db "github.com/mcctrix/ctrix-social-go-backend/internal/infrastructure/database"
	"github.com/mcctrix/ctrix-social-go-backend/internal/pkg/utils"
	"github.com/mcctrix/ctrix-social-go-backend/internal/server"
)

func main() {

	config.Load()
	CheckArgs()

	port := utils.GetEnv("PORT", "4000")

	mainRouter := server.NewServer()

	err := mainRouter.Listen(":" + port)
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
