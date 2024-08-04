package httpserv

import (
	"fmt"
	"migrate/httpserv/route"

	"github.com/witwoywhy/go-cores/apps"
	"github.com/witwoywhy/go-cores/gins"
)

func Run() {
	app := gins.New()

	app.UseMiddleware(gins.Log())
	app.UseMiddleware(gins.Error())

	route.BindMigrationRoute(app)

	app.ListenAndServe(fmt.Sprintf(":%s", apps.AppConfig.Port))
}
