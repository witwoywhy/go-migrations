package handler

import (
	"migrate/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/gins"
)

type migrationHandler struct {
	data   domain.Service[domain.Request]
	schema domain.Service[domain.Request]
}

func NewMigrationHandler(
	data domain.Service[domain.Request],
	schema domain.Service[domain.Request],
) *migrationHandler {
	return &migrationHandler{
		data:   data,
		schema: schema,
	}
}

func (h *migrationHandler) Handle(ctx *gin.Context) {
	l := gins.NewLogFromCtx(ctx)

	var request domain.Request
	var response []string
	request.Action = domain.Action(ctx.Param("action"))

	if err := ctx.BindJSON(&request); err != nil {
		ctx.Error(errs.NewBadRequestError())
		return
	}

	if domain.IsNotAction(request.Action) {
		ctx.Error(errs.NewBadRequestError())
		return
	}

	if request.Schema != nil {
		resp, err := h.schema.Execute(request, l)
		if err != nil {
			ctx.JSON(err.Status(), err)
			return
		}
		response = append(response, resp)
	}

	if request.Data != nil {
		resp, err := h.data.Execute(request, l)
		if err != nil {
			ctx.JSON(err.Status(), err)
			return
		}
		response = append(response, resp)
	}

	ctx.JSON(http.StatusOK, response)
}
