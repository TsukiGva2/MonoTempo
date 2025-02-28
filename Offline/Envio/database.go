package main

import (
	"github.com/mytempoesp/mysql-easy"
)

func (envio *Envio) SetupDatabase() (err error) {

	db, err := mysql_easy.ConfiguraDB()

	envio.db = db

	return
}

func (envio *Envio) CloseDatabase() {

	envio.db.Close()
}
