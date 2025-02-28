package main

type Atleta struct {
	Tempo      string `json:"tempo"`
	Antena     int    `json:"antena"`
	Numero     int    `json:"numero"`
	PercursoID int    `json:"percurso"`
	Staff      int    `json:"staff"`
	Check      int    `json:"checkpoint_id"`

	provaID int
}
