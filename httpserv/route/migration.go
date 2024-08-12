package route

import (
	"migrate/httpserv/handler"
	"migrate/services/data"
	"migrate/services/schema"
	"net/http"

	"github.com/witwoywhy/go-cores/gins"
)

func BindMigrationRoute(app gins.GinApps) {
	data := data.New()
	schema := schema.New()

	hdl := handler.NewMigrationHandler(data, schema)
	app.Register(
		http.MethodPost,
		"/:action",
		hdl.Handle,
	)
}
