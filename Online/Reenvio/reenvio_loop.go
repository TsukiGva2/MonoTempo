package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"
	"github.com/MyTempoESP/Reenvio/atleta"
)

const (
	REENVIO_INTERVALO     = 23 * time.Second
	TIMEOUT_MONITORAMENTO = 5 * time.Second
)

var (
	/*
		URL para subir os tempos.
	*/
	UrlTempos = fmt.Sprintf(
		"http://%s/receive/tempos", os.Getenv("MYTEMPO_API_URL"))
)

/*
	{
	    "equipamentoId": 103,
	    "provaId": 201,
	    "atletas": [
	    	{
			"antena": 1,
			"numero": 201,
			"tempo": "10:20:33.323",
			"percurso": 120
		}
	    ]
	}
*/
type AtletasForm struct {
	EquipamentoId int             `json:"equipamentoId"`
	ProvaId       int             `json:"provaId"`
	Atletas       []atleta.Atleta `json:"atletas"`
}

func (reenvio *Reenvio) Upload(atletas []atleta.Atleta) {

	if len(atletas) == 0 {

		log.Println("Ignorando lote vazio de atletas para reenvio.")

		return
	}

	if reenvio.Equip.ID == 0 {

		log.Println("Equipamento inválido! (0)")

		return
	}

	if reenvio.Equip.ProvaID == 0 {

		log.Println("Prova inválida! (0)")

		return
	}

	dados := AtletasForm{
		reenvio.Equip.ID,
		reenvio.Equip.ProvaID,
		atletas,
	}

	data, err := json.Marshal(dados)

	if err != nil {

		log.Println("Erro na conversão dos dados: ", err)

		return
	}

	/* NOTE: DEBUG */
	log.Printf("Dados para a request: %s\n", string(data))

	/*
		O erro na request não é retornado, apenas lidamos com ele.
	*/
	uploadErr := SimpleRawRequest(UrlTempos, data, "application/json")

	if uploadErr != nil {

		log.Printf("Tentativa de reenvio de um lote falhou. (erro: %s)\n", uploadErr)
	}

	return
}

/*
By Rodrigo Monteiro Junior
ter 17 set 2024 12:34:47 -03

# TentarReenvio

Aguarda por dados de reenvio e tenta o envio para a API,
em caso de erro, redireciona para o Banco de Dados.
*/
func (reenvio *Reenvio) TentarReenvio() {

	//var ( tempos []atleta.Atleta ) TODO

	/*
		Envia um lote.
	*/

	/*
		Criar o timeout para o monitoramento.
		define o tempo gasto esperando um tempo vir da lista.
	*/
	timeoutMon := time.After(
		min(TIMEOUT_MONITORAMENTO, REENVIO_INTERVALO-1),
	)

	/*
		Se houver registros, receba.
		Caso contrário, não bloqueie o código.
	*/
	select {
	//case tempos = <-reenvio.Atletas: // TODO
	//reenvio.Upload(tempos) // TODO

	case <-timeoutMon:
		log.Println("Timeout, deixando para enviar depois")
	}
}

/*
By Rodrigo Monteiro Junior
ter 17 set 2024 09:31:40 -03

# LoopReenvio

Inicia o loop de conexão com a API, reenviando atletas
dentro da queue de reenvio.
*/
func (r *Reenvio) EnviaLoop(timerEnvio *time.Ticker) {

	/*
		A cada minuto é feito o reenvio de todos
		os atletas aguardando, conforme necessidade.
	*/

	for {
		<-timerEnvio.C

		log.Println("Tentando reenvio dos atletas.")
		go r.TentarReenvio()
	}
}
