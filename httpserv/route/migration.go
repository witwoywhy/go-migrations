package route

import (
	"migrate/httpserv/handler"
	"migrate/services/migration"
	"net/http"

	"github.com/witwoywhy/go-cores/gins"
)

func BindMigrationRoute(app gins.GinApps) {
	svc := migration.New()

	hdl := handler.NewMigrationHandler(svc)
	app.Register(
		http.MethodPost,
		"/:action",
		hdl.Handle,
	)
}
