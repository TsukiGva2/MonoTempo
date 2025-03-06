package main

import (
	"log"
	"time"
)

func main() {

	var r Reenvio

	err := r.Equip.Atualiza()

	if err != nil {

		log.Fatalf("Erro atualizando equipamento: %s\n", err)
	}

	/*
		Timer para sincronizar os envios.
	*/
	timerEnvio := time.NewTicker(REENVIO_INTERVALO)

	/*
		Iniciar o loop que faz a conexão com a API e envia os atletas
		obtidos (seja por meio de não envio ou inválidos).
	*/
	r.EnviaLoop(timerEnvio)
}
