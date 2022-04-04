package main

import (
	"tf-sso/model"
	"tf-sso/router"
	"tf-sso/util/log"
)

func init() {
	log.Init()
	model.Init()
}

func main() {
	e := router.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
