package handler

import (
	"migrate/domain"
	"migrate/services/migration"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/gins"
)

type migrationHandler struct {
	service migration.Service
}

func NewMigrationHandler(service migration.Service) *migrationHandler {
	return &migrationHandler{
		service: service,
	}
}

func (h *migrationHandler) Handle(ctx *gin.Context) {
	l := gins.NewLogFromCtx(ctx)

	var request migration.Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(errs.NewBadRequestError())
		return
	}

	request.Action = domain.Action(ctx.Param("action"))

	err := h.service.Execute(request, l)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
