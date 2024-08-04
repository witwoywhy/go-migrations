package main

import (
	"migrate/httpserv"
	"migrate/infrastructure"
)

func init() {
	infrastructure.InitConfig()

}

func main() {
	infrastructure.InitLog()
	infrastructure.InitDb()
	httpserv.Run()
}
