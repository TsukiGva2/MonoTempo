package main

import "log"

/*
	"percursos": [
	        {
	            "id":                                    400 ,
	            "fimDaLargada":                    "10:55:00",
	            "horaDaLargada":                   "10:55:00",
	            "tempoMinimoParaChegada":          "10:55:00",
	            "tempoMinimoParaChegadaEmMinutos": "10:55:00"
	        },
	]
*/

type Tempo string

type Percurso struct {
	ID      int    `json:"id"`
	Desc    string `json:"descricao"`
	Inicio  Tempo  `json:"horaDaLargada"`
	Largada Tempo  `json:"fimDaLargada"`
	Chegada Tempo  `json:"tempoMinimoParaChegada"`
}

func (r *Receba) LimpaPercursosForaDaProva(provaID int) {

	if provaID == 0 {

		return
	}

	_, err := r.db.Exec(
		QUERY_LIMPA_TODOS_PERCURSOS_FORA_DA_PROVA,
		provaID,
	)

	if err != nil {
		log.Println(err)
	}
}

func (r *Receba) AtualizaPercurso(p Percurso, idProva int) (err error) {

	_, err = r.db.Exec(
		QUERY_ATUALIZA_PERCURSO,

		p.ID,
		p.Desc,
		idProva,
		p.Inicio,
		p.Chegada,
		p.Largada,
	)

	return
}
