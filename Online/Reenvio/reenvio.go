package main

import (
	"database/sql"
	//rabbit "github.com/mytempoesp/rabbit"
)

type Reenvio struct {
	//broker rabbit.Rabbit

	tempos *sql.DB

	Equip Equipamento
}
