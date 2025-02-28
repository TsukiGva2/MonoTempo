package main

import (
	"log"
	"time"
)

const (
	QUERY_ATLETAS_LARGADA = `SELECT athlete_num,
					athlete_time,
					athlete_time,
					checkpoint_id,
					antenna,
					staff,
					event_id,
					track_id
				FROM resultados_largada`

	QUERY_ATLETAS_CHEGADA = `SELECT athlete_num,
					athlete_time,
					athlete_time,
					checkpoint_id,
					antenna,
					staff,
					event_id,
					track_id
				FROM resultados_largada`
)

/*
Renomear função para inválidos.
*/

func (reenvio *Reenvio) ColetaVálidos(query string) (err error, qtd int) {

	res, err := reenvio.db.Query(query)

	if err != nil {
		log.Println("Erro checando atletas disponíveis para largada", err)

		return
	}

	defer res.Close()

	/*
		lista contendo os Atletas para o envio.
	*/
	var (
		atletas []Atleta
	)

	for res.Next() {
		var at Atleta

		err = res.Scan(
			&at.Numero,
			&at.Tempo,
			&at.Check,
			&at.Antena,
			&at.Staff,

			&at.provaID,
			&at.PercursoID,
		)

		if err != nil {
			log.Println("Erro ao escanear os atletas: ", err)

			break
		}

		atletas = append(atletas, at)

		qtd++
	}

	log.Println("Reenviando atletas")

	ReportarLoteAtletas(atletas)

	err = res.Err()

	if err != nil {
		log.Println("Erro ao escanear os atletas: ", err)

		return
	}

	/*
		Enviando atletas.
	*/
	reenvio.Atletas <- atletas

	return
}

func (reenvio *Reenvio) EnviaLoop(timerEnvio *time.Ticker) {

	go func() {
		log.Println("Aguardando percursos da prova")

		próximoEnvio := timerEnvio.C

		for {

			<-próximoEnvio

			/* TODO: wait for this */
			go func() {
				err, _ := reenvio.ColetaVálidos(QUERY_ATLETAS_LARGADA)

				if err != nil {

					log.Println("Erro ao enviar:", err)
				}

				err, _ = reenvio.ColetaVálidos(QUERY_ATLETAS_CHEGADA)

				if err != nil {

					log.Println("Erro ao enviar:", err)
				}
			}()
		}
	}()
}
