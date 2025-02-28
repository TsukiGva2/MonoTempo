package main

import (
	"log"
)

func (envio *Envio) ChecaValidade(numero int) (válido bool, err error) {

	query := `
SELECT num FROM athletes WHERE num = ?
`

	válido = false

	// reseba!!!!!!!
	res, err := envio.db.Query(query, numero)

	if err != nil {
		return
	}

	defer res.Close()

	if res.Next() {
		válido = true
	}

	err = res.Err()

	return
}

func (envio *Envio) Process() {

	tags := envio.Tags

	for {

		t := <-tags

		if t.Antena == 0 {
			/*
				Antena 0 não existe
			*/

			continue
		}

		//tempoAtleta, err := time.Parse(t.TempoFormato, t.TempoFormatado)

		/*
			lógica pra descobrir a Prova e o Percurso do atleta
			se não encontrado um deles, salvar como inválido.
		*/

		válido, err := envio.ChecaValidade(t.Epc)

		if err != nil {
			log.Printf("Erro ao obter percurso/prova do atleta %d: %+v\n", t.Epc, err)

			continue
		}

		at := Atleta{
			Antena: t.Antena,         /* Antena    */
			Numero: t.Epc,            /* Numero    */
			Staff:  t.Staff,          /* Staff     */
			Tempo:  t.TempoFormatado, /* Tempo     */
			Check:  0,                /* Check     */ //TODO
		}

		if !válido {

			envio.SalvarAtletaInvalido(&at)

			/* Atleta sem prova/percurso, ignorando */
			continue
		}

		/*
			Tentar salvar o atleta no banco de dados.
		*/

		envio.SalvarAtleta(&at)
	}
}
