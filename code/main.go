package main

import (
	"code/router"
	"code/utils"
)

func main() {
	utils.InitConfigMYSQL()
	utils.InitConfigREDIS()

	// utils.DB.AutoMigrate(&models.Contact{})
	// utils.DB.AutoMigrate(&models.Group{})

	// fmt.Println("dsdsdsd")
	r := router.Router()
	r.Run(":9090")
}
