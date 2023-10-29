package main

import (
	"go-mini-project/database"
	"go-mini-project/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.Use(middleware.CORS())
	app.Pre(middleware.RemoveTrailingSlash())

	route.NewRoute(app, db)
	app.Logger.Fatal(app.Start(":8080"))
}
