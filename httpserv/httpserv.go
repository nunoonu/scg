package httpserv

import (
	"github.com/gin-gonic/gin"
	"scg/errs"
)

func Run() {

	g := gin.New()
	g.Use(errs.ErrorHandler)
	bindCashier(g)
	bindNumber(g)
	bindHealth(g)
	g.Run(":1323")

}
