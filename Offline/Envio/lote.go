package main

import (
	"fmt"
	"log"
	"time"

	mysql "github.com/go-sql-driver/mysql"
)

const (
	LOTE_INTERVALO_VÁLIDOS   = 5 * time.Second
	LOTE_INTERVALO_INVÁLIDOS = 10 * time.Second

	ARQUIVO_DE_LOTE_VÁLIDO   = "vlote"
	ARQUIVO_DE_LOTE_INVÁLIDO = "invlote"
)

func (envio *Envio) SalvarLoteSQL(filename string, tabela string) (err error) {
	log.Println("Salvando lote SQL vindo do arquivo:", filename)

	query := fmt.Sprintf(`
LOAD DATA LOCAL INFILE '%s'
INTO TABLE %s
FIELDS TERMINATED BY ','
OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
-- IGNORE 1 LINES
(antenna, athlete_num, athlete_time, staff, checkpoint_id);
`, filename, tabela)

	_, err = envio.db.Exec(query)

	return
}

func (envio *Envio) SalvaLotes() {

	log.Println("Iniciando salvamento de lotes")

	válidos, err := ArquivoTemporário(ARQUIVO_DE_LOTE_VÁLIDO)
	inválidos, err := ArquivoTemporário(ARQUIVO_DE_LOTE_INVÁLIDO)

	if err != nil {
		log.Println("Não foi possível abrir os arquivos de lote:", err)

		return
	}

	mysql.RegisterLocalFile(válidos.Caminho)
	mysql.RegisterLocalFile(inválidos.Caminho)

	salvarAtletasVálidos := time.After(LOTE_INTERVALO_VÁLIDOS)
	salvarAtletasInválidos := time.After(LOTE_INTERVALO_INVÁLIDOS)

	go válidos.Observar()
	go inválidos.Observar()

	go func() {
		for {
			select {

			/* Receber dados de atleta e concatenar no arquivo */
			case atleta := <-envio.FilaAtletasVálidos:
				válidos.Inserir(atleta)

			case atleta := <-envio.FilaAtletasInválidos:
				inválidos.Inserir(atleta)
			}
		}
	}()

	for {
		select {

		/* Salvar dados do arquivo no banco de dados */
		case <-salvarAtletasVálidos:
			log.Println("Armazenando os tempos válidos...")

			err = envio.SalvarLoteSQL(válidos.Caminho, "athletes_times")

			if err != nil {
				log.Println("Erro ao armazenar os tempos válidos:", err)
			}

			válidos.Limpar()

			salvarAtletasVálidos = time.After(LOTE_INTERVALO_VÁLIDOS)

		case <-salvarAtletasInválidos:
			log.Println("Armazenando os tempos inválidos...")

			err = envio.SalvarLoteSQL(inválidos.Caminho, "invalidos")

			if err != nil {
				log.Println("Erro ao armazenar os tempos inválidos:", err)
			}

			inválidos.Limpar()

			salvarAtletasInválidos = time.After(LOTE_INTERVALO_INVÁLIDOS)
		}
	}
}
