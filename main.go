package main

import (
	"fmt"
	"go_todo/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Panicln("test")

	fmt.Println(models.Db)
}
