package atleta

type Atleta struct {
	Tempo  string `json:"tempo"`
	Antena int    `json:"antena"`
	Numero int    `json:"numero"`
	Staff  int    `json:"staff"`

	ProvaID    int
	PercursoID int
}
