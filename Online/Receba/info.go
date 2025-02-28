package main

import (
	"log"
	"strconv"
)

const (
	/* Actions do equipamento */

	RESTART = iota
	RESET
	FULL_RESET
	UPDATE
)

func (r *Receba) GetValidos() (validos int) {

	res, err := r.db.Query("SELECT validos FROM stats")

	if err != nil {

		return
	}

	if res.Next() {

		res.Scan(&validos)
	}

	res.Close()

	return
}

func (r *Receba) TagsTotal() (tags_total int) {

	res, err := r.db.Query("SELECT tags_total FROM stats")

	if err != nil {

		return
	}

	if res.Next() {

		res.Scan(&tags_total)
	}

	res.Close()

	return
}

func Abs(x int) int {

	if x < 0 {
		return -x
	}

	return x
}

/*
Função que envia a request para a rota de atualização
do estado de comunicação do equipamento.
Por enquanto apenas o estado da conexão com as redes
é reportado.
*/
func (r *Receba) EnviarInfo(equipID int) (err error) {

	validos := r.GetValidos()
	registros := r.TagsTotal()

	data := Form{

		"deviceId": strconv.Itoa(equipID),
		/* TODO: Redes conectadas */

		"validos":   strconv.Itoa(validos),
		"invalidos": strconv.Itoa(Abs(registros - validos)),
	}

	action, err := GetAction(r.InfoRota, data)

	if err != nil {

		log.Println("erro obtendo action:", err)

		return
	}

	if action == 0 {

		return
	}

	log.Printf("Obteve ação: %d\n", action)

	IgnorarForeignKey(r.db)

	switch action {
	case FORMATAR:
		_, err = r.db.Exec(`
			TRUNCATE checkpoints;
			TRUNCATE equipamento;
			TRUNCATE event_data;
			TRUNCATE athletes;
			TRUNCATE tracks;
			TRUNCATE stats
		`)

		fallthrough
	case EXCLUIR:
		_, err = r.db.Exec(`
			TRUNCATE invalidos;
			TRUNCATE athletes_times
		`)
	case REINICIAR:
		// restart
	}

	AceitarForeignKey(r.db)

	return
}
