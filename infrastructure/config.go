package infrastructure

import (
	"github.com/witwoywhy/go-cores/apps"
	"github.com/witwoywhy/go-cores/vipers"
)

func InitConfig() {
	vipers.Init()
	apps.InitAppConfig()
}
