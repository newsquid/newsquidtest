package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/newsquid/mscore"
	"github.com/newsquid/newsquidtest/assignment2/backend/api"
)

/*
Application entry point
*/
func main() {
	m := martini.New()

	//Set middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"OPTIONS", "HEAD", "POST", "GET", "PUT"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	//Create the router
	r := martini.NewRouter()

	//Options matches all and sends okay
	r.Options(".*", func() (int, string) {
		return 200, "ok"
	})

	api.SetupTodoRoutes(r)
	api.SetupCommentRoutes(r)

	mscore.StartServer(m, r)
	fmt.Println("Started NSQ Test service")
}
