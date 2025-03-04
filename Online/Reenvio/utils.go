package main

import (
	"log"
)

/*
By Rodrigo Monteiro Junior
qui 12 set 2024 09:47:17 -03

# ReportarLoteAtletas

NOTE: DEBUG

função para reportar um lote de atletas
*/

func ReportarLoteAtletas(lote []Atleta) {
	for i, atleta := range lote {
		log.Printf("ATLETA ENVIADO: %+v\n", atleta)

		if i >= 3 {
			log.Println("Output cortado por brevidade")

			return
		}
	}
}
