package main

import (
	"fmt"
	"log"
	"time"
)

type Equipamento struct {
	Nome string

	ID      int
	ProvaID int
}

const (
	ATUALIZA_EQUIP_INTERVALO = 1 * time.Minute
)

func (reenvio *Reenvio) AtualizaEquip() (err error) {

	query := `SELECT idequip, modelo, event_id FROM equipamento WHERE 1;`

	res, err := reenvio.db.Query(query)

	if err != nil {
		return
	}

	defer res.Close()

	if !res.Next() {
		err = fmt.Errorf("Dados do dispositivo não encontrados.")

		return
	}

	err = res.Scan(
		&reenvio.Equip.ID,
		&reenvio.Equip.Nome,
		&reenvio.Equip.ProvaID,
	)

	if err != nil {
		return
	}

	err = res.Err()

	return
}

func (reenvio *Reenvio) AtualizarEquip() {

	t := time.NewTicker(ATUALIZA_EQUIP_INTERVALO)

	for {

		<-t.C

		log.Println("Atualizando dados do equipamento")

		err := reenvio.AtualizaEquip()

		if err != nil {
			log.Println("Não foi possível atualizar o equipamento.", err)

			continue
		}
	}
}
