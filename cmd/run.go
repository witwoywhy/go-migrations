package cmd

import (
	"flag"
	"migrate/domain"
	"migrate/services/data"
	"migrate/services/schema"

	"github.com/google/uuid"
	"github.com/witwoywhy/go-cores/apps"
	"github.com/witwoywhy/go-cores/errs"
	"github.com/witwoywhy/go-cores/logs"
)

func Run() {
	l := logs.New(map[string]any{
		apps.TraceID: uuid.NewString(),
		apps.SpanID:  uuid.NewString(),
	})

	action := flag.String("action", "up", "for migration 'up' or 'down'")
	migrate := flag.String("migrate", "schema", "for migration 'schema' or 'data'")
	forceVersion := flag.Int("fvd", 0, "for force migration  version, should more than 0")

	flag.Parse()

	if domain.IsNotAction(domain.Action(*action)) {
		l.Error("wrong action")
		return
	}

	m := domain.MigrateType(*migrate)

	if domain.IsNotMigrateType(m) {
		l.Error("wrong migrate")
		return
	}

	var response string
	var err errs.Error

	switch m {
	case domain.Schema:
		schema := schema.New()
		response, err = schema.Execute(domain.Request{
			Action: domain.Action(*action),
			Schema: &domain.Migrate{
				ForceVersion: *forceVersion,
			},
			Data: &domain.Migrate{},
		}, l)
	case domain.Data:
		data := data.New()
		response, err = data.Execute(domain.Request{
			Action: domain.Action(*action),
			Schema: &domain.Migrate{},
			Data: &domain.Migrate{
				ForceVersion: *forceVersion,
			},
		}, l)
	}

	if err != nil {
		l.Error(err)
	} else {
		l.Info(response)
	}
}
