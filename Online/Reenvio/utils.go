package main

import (
	"encoding/json"

	"log"
)

/*
# AtletasFromJSON

recebe uma body em json "marshal-izado", no formato:

```

	[]byte(
		`[
		   {
		      tempo         ,
		      antena        ,
		      numero        ,
		      percurso      ,
		      staff         ,
		      checkpoint_id ,
		   }
		]`,
	)

```

e converte para []Atleta, cuja expansão é algo parecido com:

```

	[]struct{
		Tempo      string `json:"tempo"`
		Antena     int    `json:"antena"`
		Numero     int    `json:"numero"`
		PercursoID int    `json:"percurso"`
		Staff      int    `json:"staff"`
		Check      int    `json:"checkpoint_id"`

		...
	}

```
*/
func AtletasFromJSON(body []byte) (atletas []Atleta, err error) {

	err = json.Unmarshal(body, &atletas)

	return
}

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
