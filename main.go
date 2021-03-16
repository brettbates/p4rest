package main

import (
	"log"

	"github.com/brettbates/p4rest/common"
	"github.com/brettbates/p4rest/config"
	routes "github.com/brettbates/p4rest/routes"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var c config.Config
	err := envconfig.Process("p4rest", &c)
	if err != nil {
		log.Printf("Failed to parse config %s\n", err)
	}

	r := gin.Default()
	r.Use(common.P4Connecter(c))
	routes.Initialize(r)
	r.RunTLS(":443", "mykey.crt", "mykey.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
