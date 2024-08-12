package main

import (
	"migrate/infrastructure"
)

func init() {
	infrastructure.InitConfig()

}

func main() {
	infrastructure.InitLog()
	infrastructure.InitDb()
	// httpserv.Run()
	// cmd.Run()
}
