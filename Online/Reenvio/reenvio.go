package main

import (
	"database/sql"
	//rabbit "github.com/mytempoesp/rabbit"
)

type Reenvio struct {
	//broker rabbit.Rabbit

	tempos *sql.DB

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
