package main

import (
	"code/router"
	"code/utils"
)

func main() {
	utils.InitConfigMYSQL()
	// fmt.Println("dsdsdsd")
	r := router.Router()
	r.Run(":9090")
}
