package main

import (

	"github.com/labstack/echo"
	"keymanager/controller"
	"os"
	"keymanager/configuration"
	"github.com/labstack/echo/middleware"
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


	r := e.Group("/api")

	// Configure middleware with the custom claims type
	//TODO: use root of trust here
	config := middleware.JWTConfig{
		Claims:     &controller.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))

	r.POST("/createuser",controller.CreateUser)
	r.POST("/deleteuser",controller.DeleteUser)
	r.PATCH("/updateuser",controller.UpdateUser)
	r.GET("/listusers",controller.ListUsers)


	r.POST("/createkey",controller.CreateKey)

	e.Logger.Fatal(e.Start(":1323"))

}
