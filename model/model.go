package model

import (
	"tf-sso/util/log"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func Init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/tf_sso")
	if err != nil {
		log.Logger.Error().Msg(err.Error())
		return
	}

	err = Engine.Ping()
	if err != nil {
		log.Logger.Error().Msg(err.Error())
		return
	}

	log.Logger.Info().Msg("Database connect at: localhost:3306")

	createTable()
}

func createTable() {
	Engine.Sync2(new(User))
}
