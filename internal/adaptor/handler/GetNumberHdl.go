package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scg/internal/core/domain"
)

type getNumberHdl struct {
	svc domain.GetNumberSvc
}

func NewGetNumberHdl(svc domain.GetNumberSvc) *getNumberHdl {
	return &getNumberHdl{svc}
}

func (hdl *getNumberHdl) Handle(g *gin.Context) {

	var req domain.GetNumberReq
	res, err := hdl.svc.Execute(req)
	if err != nil {
		g.Error(err)
		return
	}
	g.JSON(http.StatusOK, &res)

}
