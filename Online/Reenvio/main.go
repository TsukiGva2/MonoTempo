package main

import (
	"log"

	"time"
)

func main() {

	/*
		Cadastrar canal de recebimento do rabbitmq para
		receber atletas não enviados com sucesso pelo Envio.
	*/

	var r Reenvio

	err := r.ConfiguraDB()

	if err != nil {
		log.Fatalf("Erro no banco de dados: %s\n", err)
	}

	r.Atletas = make(chan []Atleta)

	/*
		Timer para sincronizar os envios.
	*/
	timerEnvio := time.NewTicker(REENVIO_INTERVALO)

	r.EnviaLoop(timerEnvio)

	/*
		Iniciar o loop que faz a conexão com a API e envia os atletas
		obtidos (seja por meio de não envio ou inválidos).
	*/
	r.LoopReenvio(timerEnvio)
}
