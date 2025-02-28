package main

import (
	"log"
	"strconv"
)

type Staff struct {
	ID   int
	Nome string
}

func (r *Receba) BuscaStaff(idProva int) (staffs []Staff, err error) {

	data := Form{
		"idProva": strconv.Itoa(idProva),
	}

	err = JSONRequest(r.StaffRota, data, &staffs)

	return
}

/*
Rodrigo Monteiro Junior
ter 10 set 2024 16:11:24 -03

Apaga todos os staffs fora da prova especificada.

  - Caso a prova seja vazia, nada é feito.
  - Use essa função com cuidado, funcionalidade não foi testada.
*/
func (r *Receba) LimpaStaffsForaDaProva(provaID int) {

	if provaID == 0 {

		return
	}

	r.db.Exec(
		QUERY_LIMPA_TODOS_STAFFS_FORA_DA_PROVA,
		provaID,
	)
}

/*
Deprecated: É preferível o uso da função LimpaStaffsForaDaProva para apagar condicionalmente.

XXX: maneira equivocada de limpar todos os staffs.
O uso dessa query deve ser pensado, pois apaga todos os registros da tabela de staffs.
*/
func (r *Receba) LimpaStaff() {

	r.db.Exec(QUERY_LIMPA_TODOS_STAFFS)
}

func (r *Receba) AtualizaStaff(staffs []Staff, idProva int) (err error) {

	for _, staff := range staffs {
		_, err := r.db.Exec(
			QUERY_ATUALIZA_STAFF,

			staff.ID,
			idProva,
			staff.Nome,
		)

		if err != nil {
			log.Printf("erro no Sql atualizando os staffs %+v\n", err)
		}
	}

	return
}
