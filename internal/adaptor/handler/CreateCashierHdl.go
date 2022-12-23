package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scg/internal/core/domain"
)

type createCashierHdl struct {
	svc domain.CreateCashierSvc
}

func NewCreateCashierHdl(svc domain.CreateCashierSvc) *createCashierHdl {
	return &createCashierHdl{svc}
}

func (hdl *createCashierHdl) Handle(g *gin.Context) {

	var req domain.CreateCashierReq

	if err := g.ShouldBindJSON(&req); err != nil {
		g.Error(err)
		return
	}
	res, err := hdl.svc.Execute(req)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(http.StatusOK, &res)

}
