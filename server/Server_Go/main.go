package main

import (
	"first-api/Config"
	"first-api/Models"
	"first-api/Routes"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Worker{})

	r := Routes.SetupRouter()

	/* Cors default * */
	r.Use(cors.Default())

	//Running
	r.Run(":3040")
}
