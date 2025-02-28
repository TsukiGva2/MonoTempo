package main

import (
	"encoding/json"
	"os"
	"time"

	"fmt"
	"log"

	"sync"
)

const (
	REENVIO_INTERVALO = 23 * time.Second

	ARMAZENA_ATLETA_QUERY = `
	INSERT INTO unsent_athletes(
	    antenna,
	    checkpoint_id,
	    athlete_num,
	    athlete_time,
	    staff,
	    percurso
	)
	VALUES (?, ?, ?, ?, ?, ?)
`

	TIMEOUT_MONITORAMENTO = 5 * time.Second
)

var (
	/*
		URL para subir os tempos.
	*/
	UrlTempos = fmt.Sprintf("http://%s/receive/tempos", os.Getenv("MYTEMPO_API_URL"))

	dbMx sync.Mutex
)

// DEPRECATED: this is useless
func (reenvio *Reenvio) ArmazenarAtletas(atletas []Atleta) {

	dbMx.Lock()
	defer dbMx.Unlock()

	for _, atleta := range atletas {

		_, err := reenvio.db.Exec(
			ARMAZENA_ATLETA_QUERY,

			atleta.Antena,
			atleta.Check,
			atleta.Numero,
			atleta.Tempo,
			atleta.Staff,
			atleta.PercursoID,
		)

		if err != nil {

			// XXX: Remover logs excessivos
			log.Printf("Falha no envio e na inserção do atleta %+v, com erro %s\n", atleta, err)
		}
	}
}

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
	EquipamentoId int      `json:"equipamentoId"`
	ProvaId       int      `json:"provaId"`
	Atletas       []Atleta `json:"atletas"`
}

func (reenvio *Reenvio) Upload(atletas []Atleta) {

	if len(atletas) == 0 {

		log.Println("Ignorando lote vazio de atletas para reenvio.")

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

		log.Printf("Tentativa de reenvio de um lote falhou. Armazenando no banco de dados (erro: %s)\n", uploadErr)

		/*
			Erro no envio dos atletas, inserir no banco de dados
			TODO.
		*/

		//reenvio.ArmazenarAtletas(atletas)
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

	var (
		inválidos []Atleta
	)

	/*
		Envia um lote de cada um (Inválidos e Não Enviados [Monitoramento]).
	*/

	/*
		Criar o timeout para o monitoramento.

		define o tempo gasto esperando um tempo vir da lista.
	*/
	timeoutMon := time.After(
		min(TIMEOUT_MONITORAMENTO, REENVIO_INTERVALO-1),
	)

	/*
		Se houver registros inválidos, receba.
		Caso contrário, não bloqueie o código.
	*/

	select {
	case inválidos = <-reenvio.Atletas:
		reenvio.Upload(inválidos)

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
func (r *Reenvio) LoopReenvio(timerEnvio *time.Ticker) {

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
