package migration

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

func New() Service {
	return &service{}
}

func (s *service) Execute(request Request, l logger.Logger) errs.Error {
	if err := request.Validate(l); err != nil {
		return err
	}

	schemaMysql, err := migrate.NewWithDatabaseInstance(
		infrastructure.Schema,
		dbs.Mysql,
		infrastructure.SchemaMysql,
	)
	if err != nil {
		l.Errorf("failed to new instance schema mysql: %v", err)
		return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
	}

	schemaPg, err := migrate.NewWithDatabaseInstance(
		infrastructure.Schema,
		dbs.Postgres,
		infrastructure.SchemaPg,
	)
	if err != nil {
		l.Errorf("failed to new instance schema pg: %v", err)
		return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
	}

	dataMysql, err := migrate.NewWithDatabaseInstance(
		infrastructure.Data,
		dbs.Mysql,
		infrastructure.DataMysql,
	)
	if err != nil {
		l.Errorf("failed to new instance data mysql: %v", err)
		return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
	}

	dataPg, err := migrate.NewWithDatabaseInstance(
		infrastructure.Data,
		dbs.Postgres,
		infrastructure.DataPg,
	)
	if err != nil {
		l.Errorf("failed to new instance data pg: %v", err)
		return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
	}

	if request.Schema != nil && request.Schema.ForceVersion > 0 {
		if err := schemaMysql.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("mysql failed to migrate schema force version: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := schemaPg.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("pg failed to migrate schema force version: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}

	if request.Data != nil && request.Data.ForceVersion > 0 {
		if err := dataMysql.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("mysql failed to migrate data force version: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := dataPg.Force(request.Schema.ForceVersion); err != nil {
			l.Errorf("pg failed to migrate data force version: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
	}

	switch request.Action {
	case domain.Up:
		if err := schemaMysql.Up(); err != nil {
			l.Errorf("mysql failed to migrate schema up: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := schemaPg.Up(); err != nil {
			l.Errorf("pg failed to migrate schema up: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}

		if err := dataMysql.Up(); err != nil {
			l.Errorf("mysql failed to migrate data up: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := dataPg.Up(); err != nil {
			l.Errorf("pg failed to migrate data up: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}

		return nil
	case domain.Down:
		if err := schemaMysql.Down(); err != nil {
			l.Errorf("mysql failed to migrate schema down: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := schemaPg.Down(); err != nil {
			l.Errorf("pg failed to migrate schema down: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}

		if err := dataMysql.Down(); err != nil {
			l.Errorf("mysql failed to migrate data down: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "mysql")
		}

		if err := dataPg.Down(); err != nil {
			l.Errorf("pg failed to migrate data down: %v", err)
			return errs.NewCustom(http.StatusInternalServerError, errs.Err50002, err.Error(), "pg")
		}
		return nil
	}

	return errs.NewBadRequestError()
}
