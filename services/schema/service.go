package schema

import (
	"migrate/domain"
	"migrate/infrastructure"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/witwoywhy/go-cores/dbs"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logger"
)

type service struct{}

func New() domain.Service[domain.Request] {
	return &service{}
}

func (s *service) Execute(request domain.Request, l logger.Logger) (string, errs.Error) {
	mysql, err := migrate.NewWithDatabaseInstance(
		infrastructure.Schema,
		dbs.Mysql,
		infrastructure.SchemaMysql,
	)
	if err != nil {
		l.Errorf("failed to new instance schema mysql: %v", err)
		return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
	}

	pg, err := migrate.NewWithDatabaseInstance(
		infrastructure.Schema,
		dbs.Postgres,
		infrastructure.SchemaPg,
	)
	if err != nil {
		l.Errorf("failed to new instance schema pg: %v", err)
		return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
	}

	if request.Schema.ForceVersion > 0 {
		if err := mysql.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("mysql failed to migrate schema force version: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := pg.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("pg failed to migrate schema force version: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}
	switch request.Action {
	case domain.Up:
		if err := mysql.Up(); err != nil {
			l.Errorf("mysql failed to migrate schema up: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := pg.Up(); err != nil {
			l.Errorf("pg failed to migrate schema up: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}

	case domain.Down:
		if err := mysql.Down(); err != nil {
			l.Errorf("mysql failed to migrate schema down: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := pg.Down(); err != nil {
			l.Errorf("pg failed to migrate schema down: %v", err)
			return "", errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}

	return "Migration Schema Success", nil
}
