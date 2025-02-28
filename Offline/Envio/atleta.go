package main

import (
	"fmt"
)

type Atleta struct {
	Tempo      string `json:"tempo"`
	Antena     int    `json:"antena"`
	Numero     int    `json:"numero"`
	PercursoID int    `json:"percurso"`
	Staff      int    `json:"staff"`
	Check      int    `json:"checkpoint_id"`
}

func (envio *Envio) SalvarAtleta(a *Atleta) {

	dados := fmt.Sprintf("%d,%d,%s,%d,%d", a.Antena, a.Numero, a.Tempo, a.Staff, a.Check)
	envio.FilaAtletasVálidos <- dados
}

func (envio *Envio) SalvarAtletaInvalido(a *Atleta) {

	dados := fmt.Sprintf("%d,%d,%s,%d,%d", a.Antena, a.Numero, a.Tempo, a.Staff, a.Check)
	envio.FilaAtletasInválidos <- dados
}
