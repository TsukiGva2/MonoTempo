package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func (r *Reenvio) ConfiguraDB(temposDB string) (err error) {

	tempos, err := sql.Open("sqlite", "/var/monotempo-data/"+temposDB)

	if err != nil {

		return
	}

	r.tempos = tempos

	return
}
