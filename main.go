package main

import (
	routes "github.com/brettbates/p4rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.Initialize(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
