package handlers

import (
	"net/http"
	"time"

	p4 "github.com/brettbates/p4go"
	"github.com/gin-gonic/gin"
)

// Fixes returns p4 fixes for a given id
func Fixes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"Fixes": []p4.Fix{
		{
			Code:   "code",
			Change: 0,
			Client: "client",
			Date:   time.Now(),
			Job:    "job",
			Status: "status",
			User:   "user"}}})
}
