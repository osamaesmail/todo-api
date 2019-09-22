package main

import (
	"github.com/labstack/echo"
	"log"
	"todo-api/db"
)

func main()  {
	log.Println("Starting server..")
	db.Init()
	e := echo.New()
	initRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}