package main

import (
	"fmt"
	"os"

	usr "github.com/civitops/Ecommercify/user/implementation/user"
	"github.com/civitops/Ecommercify/user/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// fetchs configuration
	cfg, err := config.LoadConfig("../../../")
	if err != nil {
		fmt.Printf("failed to load config: %s", err.Error())
		os.Exit(1)
	}
	// connecting postgres DB through Go-ORM
	pgConn, err := gorm.Open(postgres.Open(cfg.DatabaseURI), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	pgConn.AutoMigrate(&usr.Entity{})
}
