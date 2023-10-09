package main

import (
	"github.com/elysiamori/assignment3-09/config"
	"github.com/elysiamori/assignment3-09/controllers"
	"github.com/elysiamori/assignment3-09/routers"
)

/*
NAME : VALDY RAMADHAN
GLNG-KS08-09
*/
func main() {
	db := config.DBConn()
	go controllers.UpdaterData(db)

	r := routers.StartRouter()
	r.Run(":3000")

}
