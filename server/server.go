package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igor570/eventexplorer/db"
	"github.com/igor570/eventexplorer/routes"
)

func main() {
	db.InitDB()
	app := gin.Default()

	routes.RegisterRoutes(app)

	//routes

	app.Run(":3100") // localhost:3100
}
