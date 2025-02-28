package main

func (r *Receba) Limpar24h() {

	r.db.Exec(
		QUERY_LIMPAR_INVALIDOS,
	)
	r.db.Exec(
		QUERY_LIMPAR_BACKUP,
	)
}
