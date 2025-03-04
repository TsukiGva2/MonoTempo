package main

import (
	"log"
	"os"
	"time"
)

func main() {

	var r Reenvio

	err := r.ConfiguraDB(os.Args[1]) // tempos db

	if err != nil {

		log.Fatalf("Erro no banco de dados: %s\n", err)
	}

	err = r.Equip.Atualiza()

	if err != nil {

		log.Fatalf("Erro atualizando equipamento: %s\n", err)
	}

	r.Atletas = make(chan []Atleta)

	/*
		Timer para sincronizar os envios.
	*/
	timerEnvio := time.NewTicker(REENVIO_INTERVALO)

	r.PreparaLoop(timerEnvio)

	/*
		Iniciar o loop que faz a conexão com a API e envia os atletas
		obtidos (seja por meio de não envio ou inválidos).
	*/
	r.EnviaLoop(timerEnvio)
}
