package httpserv

import (
	"github.com/gin-gonic/gin"
	"scg/fnd"
	"scg/internal/adaptor/handler"
	"scg/internal/core/service"
)

func bindCashier(g *gin.Engine) {

	svc := service.NewCreateCashierSvc()
	hdl := handler.NewCreateCashierHdl(svc)
	g.POST("/cashier", hdl.Handle)

}

func bindNumber(g *gin.Engine) {

	svc := service.NewGetNumberSvc()
	hdl := handler.NewGetNumberHdl(svc)
	g.GET("/number", hdl.Handle)

}

func bindHealth(a *gin.Engine) {
	a.GET("/health", fnd.HealthHandler)
}
