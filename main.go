package main

import (
	"github.com/gucarletto/go-gin-rest-boilerplate/db"
	"github.com/gucarletto/go-gin-rest-boilerplate/server"
)

func main() {
	db.Migrate()
	server.Start()
}
