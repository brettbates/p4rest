package handlers

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	p4 "github.com/brettbates/p4go"
	"github.com/brettbates/p4rest/common"
	"github.com/gin-gonic/gin"
)

// Fixes returns p4 fixes for a given id
func Fixes(ctx *gin.Context) {
	cl := ctx.DefaultQuery("change", "1")
	p4r := ctx.MustGet("p4c").(common.P4Runner)

	fixes, err := p4.RunFixes(p4r, []string{fmt.Sprintf("//...@%s,@%s", cl, cl)})
	if err != nil {
		// TODO Log the erorr and request
		log.Printf("Failed to retrieve fixes %s\n", err)
	}
	ctx.JSON(http.StatusOK, gin.H{"Fixes": fixes})
}
