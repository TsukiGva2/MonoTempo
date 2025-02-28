package main

import (
	"database/sql"
)

const (
	QUERY_LIMPAR_INVALIDOS_IMEDIATO = `
	DELETE FROM invalidos;
	`

	QUERY_LIMPAR_BACKUP_IMEDIATO = `
	DELETE FROM athletes_times;
	`

	QUERY_LIMPAR_EQUIP = `
	DELETE FROM equipamento;
	`

	QUERY_LIMPAR_INVALIDOS = `
	DELETE FROM invalidos WHERE timestp <= NOW() - INTERVAL 2 DAY
	`

	QUERY_LIMPAR_BACKUP = `
	DELETE FROM athletes_times WHERE timestp <= NOW() - INTERVAL 2 DAY
	`

	/*
		Deprecated: Remover essa query.
		XXX: Encontrar uma forma mais segura de substituir os staffs
	*/
	QUERY_LIMPA_TODOS_STAFFS = `
	DELETE FROM staff
	`

	/*
		Deprecated: Remover essa query.
		XXX: Encontrar uma forma mais segura de substituir os atletas
	*/
	QUERY_LIMPA_TODOS_ATLETAS = `
	DELETE FROM athletes
	`

	QUERY_LIMPA_TODOS_STAFFS_FORA_DA_PROVA = `
	DELETE FROM staffs WHERE event_id <> ?
	`

	QUERY_LIMPA_TODOS_ATLETAS_FORA_DA_PROVA = `
	DELETE FROM athletes WHERE event_id <> ?
	`

	QUERY_LIMPA_TODOS_PERCURSOS_FORA_DA_PROVA = `
	DELETE FROM tracks WHERE event_id <> ?
	`

	/* FIXME: o Equipamento padrão é denominado 'X'. */
	QUERY_ATUALIZA_EQUIP = `
	INSERT INTO equipamento(
	    id,
	    idequip,
	    modelo,
	    event_id,
	    checkpoint_id
	)
	VALUES(
		1,
		0,
		"",
		0,
		0
	)
	ON DUPLICATE KEY
	UPDATE
	    idequip = ?,
	    modelo = ?,
	    event_id = ?,
	    checkpoint_id = ?
	`

	QUERY_ATUALIZA_PROVA = `
	REPLACE INTO event_data (
		id,
		event_date,
		event_title
	)
	VALUES (?, ?, ?) 
	`

	QUERY_ATUALIZA_PERCURSO = `
	REPLACE INTO tracks (
	    id,
	    race_description,
	    event_id,
	    inicio,
	    chegada,
	    largada
	)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	QUERY_ATUALIZA_ATLETA = `
	REPLACE INTO athletes (
	    name,
	    sex,
	    team,
	    city,
	    num,
	    event_id,
	    track_id
	)
	VALUES(?, ?, ?, ?, ?, ?, ?);
	`

	QUERY_ATUALIZA_STAFF = `
	REPLACE INTO staffs (
	    id,
	    event_id,
	    nome
	)
	VALUES(?, ?, ?);
	`
)

func IgnorarForeignKey(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS=0")
}

func AceitarForeignKey(db *sql.DB) {
	db.Exec("SET FOREIGN_KEY_CHECKS=1")
}
