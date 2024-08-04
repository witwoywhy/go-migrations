package migration

import (
	"migrate/domain"

	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

type Service interface {
	Execute(request Request, l logger.Logger) errs.Error
}

type Request struct {
	Action domain.Action
	Schema *domain.Migrate `json:"schema"`
	Data   *domain.Migrate `json:"data"`
}
