package main

import (

	"github.com/labstack/echo"
	"keymanager/controller"
	"os"
	"keymanager/configuration"
)

func main() {

	os.Setenv("DBNAME","Keymanager")
	os.Setenv("DBUSR","root")
	os.Setenv("DBPASSWORD","root")
	os.Setenv("DBHOST","0.0.0.0")
	os.Setenv("DBPORT","3456")


	configuration.InitializeConfiguration()

	e := echo.New()
	e.POST("/login",controller.LoginApi)
	e.Logger.Fatal(e.Start(":1323"))
}
