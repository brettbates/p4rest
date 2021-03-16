package common

import (
	p4 "github.com/brettbates/p4go"
	"github.com/brettbates/p4rest/config"
	"github.com/gin-gonic/gin"
)

// P4Connecter adds a p4c connection to the context
func P4Connecter(c config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p4c := NewP4CParams(c)
		ctx.Set("p4c", p4c)
	}
}

// P4Runner is an interface for testing without calling p4
type P4Runner interface {
	Run([]string) ([]map[interface{}]interface{}, error)
}

// P4C Is a wrapper around the p4 connection
type P4C struct {
	p4.P4
}

// NewP4C connects to p4 and returns a P4C wrapper
func NewP4C() *P4C {
	return &P4C{P4: *p4.NewP4()}
}

// NewP4CParams TODO This needs to read from .p4config files
func NewP4CParams(c config.Config) *P4C {
	return &P4C{P4: *p4.NewP4Params(c.P4Port, c.P4User, c.P4Client)}
}
