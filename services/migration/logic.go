package migration

import (
	"migrate/domain"

	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

func (r *Request) Validate(l logger.Logger) errs.Error {
	if r.Schema == nil && r.Data == nil {
		l.Error("request.schema and request.data is null")
		return errs.NewBadRequestError()
	}

	if domain.IsNotAction(r.Action) {
		l.Error("is not action")
		return errs.NewBadRequestError()
	}

	return nil
}
