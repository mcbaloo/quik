package main

import (
	"fmt"

	"quik-assessment/Config"
	"quik-assessment/Models"
	"quik-assessment/Routes"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	//Run Migration
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Wallet{}, &Models.Transaction{})
	Config.DB.Model(&Models.Transaction{}).AddForeignKey("wallet_id", "wallet(id)", "RESTRICT", "RESTRICT")
	r := Routes.SetupRouter()
	r.Run()
}
