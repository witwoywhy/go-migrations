package main

import (
	"migrate/infrastructure"
	"sync"
)

func init() {
	infrastructure.InitConfig()

}

func combineRun() {
	var wg sync.WaitGroup
	wg.Add(1)
}

func main() {
	infrastructure.InitLog()
	infrastructure.InitDb()
	// httpserv.Run()
	// cmd.Run()
}
