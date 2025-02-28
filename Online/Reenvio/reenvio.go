package main

import (
	"database/sql"

	mysql "github.com/mytempoesp/mysql-easy"
	rabbit "github.com/mytempoesp/rabbit"
)

type Reenvio struct {
	broker rabbit.Rabbit
	db     *sql.DB

	Equip Equipamento

	/*
		Pode ser inserido no canal `Atletas` um tempo
		que se encaixe em qualquer um dos seguintes critérios:

			- É considerado válido, após reavaliação, nas restrições da prova
			  e dentro das regras gerais do funcionamento do envio*;

			      ( * NOTE: Regras do envio

			      considerando um tempo t:
			              t >= Início
			              t <= Largada

			                ou

			              t > Largada
			      )

			- Se encontra dentro da tabela `unsent_athletes`;
			- É um tempo de Largada.
	*/
	Atletas chan []Atleta
}

func (reenvio *Reenvio) ConfiguraDB() (err error) {

	db, err := mysql.ConfiguraDB()

	reenvio.db = db

	return
}
