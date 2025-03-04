package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func (r *Receba) ConfiguraDB() (err error) {

	db, err := sql.Open("sqlite", "/var/monotempo-data/equipamento.sqlite")

	if err != nil {

		return
	}

	r.db = db

	return
}
