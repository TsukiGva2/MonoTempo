package main

import (
	//"log"
	"strconv"
)

/*
{
	"numero": 420,
	"provaId": 201,
	"percursoId": 401,
	"nome": "DALVA",
	"equipe": "JATO"
}
*/

type Atleta struct {
	Nome       string `json:"nome"`
	Sexo       string `json:"sexo"`
	Equipe     string `json:"equipe"`
	Cidade     string `json:"cidade"`
	Numero     int    `json:"numero"`
	ProvaID    int    `json:"provaId"`
	PercursoID int    `json:"percursoId"`
}

func (r *Receba) BuscaAtletas(idProva int) (atletas []Atleta, err error) {

	data := Form{
		"idProva": strconv.Itoa(idProva),
	}

	err = JSONRequest(r.AtletasRota, data, &atletas)

	return
}

/*
Deprecated: É preferível o uso da função LimpaAtletasForaDaProva para apagar condicionalmente.

XXX: maneira equivocada de limpar todos os atletas.
O uso dessa query deve ser pensado, pois apaga todos os registros da tabela de atletas.
*/
func (r *Receba) LimpaAtletas() {

	r.db.Exec(QUERY_LIMPA_TODOS_ATLETAS)
}

/*
Rodrigo Monteiro Junior
ter 10 set 2024 14:38:50 -03

Apaga todos os atletas fora da prova especificada.

  - Caso a prova seja vazia, nada é feito.
  - Use essa função com cuidado, funcionalidade não foi testada.

TODO: Testes
*/
func (r *Receba) LimpaAtletasForaDaProva(provaID int) {

	if provaID == 0 {

		return
	}

	r.db.Exec(
		QUERY_LIMPA_TODOS_ATLETAS_FORA_DA_PROVA,
		provaID,
	)
}

func (r *Receba) AtualizaAtletas(atletas []Atleta) (err error) {

	for _, a := range atletas {
		_, err = r.db.Exec(
			QUERY_ATUALIZA_ATLETA,

			a.Nome,
			a.Sexo,
			a.Equipe,
			a.Cidade, /* city */
			a.Numero,
			a.ProvaID,
			a.PercursoID,
		)

		//log.Println(a.Sexo)
	}

	return
}
