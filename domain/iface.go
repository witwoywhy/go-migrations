package domain

import (
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

type Service[Request any] interface {
	Execute(request Request, l logger.Logger) (string, errs.Error)
}
