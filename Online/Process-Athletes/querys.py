percursosSQL = "SELECT * FROM tracks"
#       LARGADA
"""
		Seleciona:
			Numero
			Tempo  <- maior possível
			Check
			Antena
			Staff

			provaID
			PercursoID

		Na tabela `athletes_times`

		com as tabelas athletes ( para obter o percurso do atleta )
		e a tabela tracks       ( o percurso para obter a largada )

		filtro:
			Tempo > tracks.Largada

		Os `group by` são necessários para que se possa
		obter o maior tempo possível. Agrupando cada tempo em
		seu respectivo atleta.
"""
largadaSQL = """
                INSERT IGNORE INTO resultados_largada (
                    athlete_num,
                    athlete_time,
                    checkpoint_id,
                    antenna,
                    staff,
                    event_id,
                    track_id
                )
                SELECT
                    athlete_num,
                    MAX(athlete_time) AS athlete_time,
                    checkpoint_id,
                    antenna,
                    staff,
                    tracks.event_id,
                    tracks.id
                FROM
                    athletes_times
                JOIN athletes ON athlete_num = athletes.num
                JOIN tracks ON tracks.id = athletes.track_id
                WHERE
                    athlete_time > tracks.inicio AND
                    athlete_time < tracks.largada AND
                    tracks.id = {tid} AND
                    staff = 0
                GROUP BY
                    athlete_num, checkpoint_id, antenna, staff, tracks.event_id, tracks.id
"""

chegadaSQL = """
                INSERT IGNORE INTO resultados_chegada (
                    athlete_num,
                    athlete_time,
                    checkpoint_id,
                    antenna,
                    staff,
                    event_id,
                    track_id
                )
                SELECT
                    athlete_num,
                    MIN(athlete_time) AS athlete_time,
                    checkpoint_id,
                    antenna,
                    staff,
                    tracks.event_id,
                    tracks.id
                FROM
                    athletes_times
                JOIN athletes ON athlete_num = athletes.num
                JOIN tracks ON tracks.id = athletes.track_id
                WHERE
                    athlete_time > tracks.largada AND
                    athlete_time > tracks.chegada AND
                    tracks.id = {tid} AND
                    staff = 0
                GROUP BY
                    athlete_num, checkpoint_id, antenna, staff, tracks.event_id, tracks.id;

"""


def largada_por_percurso(id_valor):
    return largadaSQL.format(tid=id_valor)


def chegada_por_percurso(id_valor):
    return chegadaSQL.format(tid=id_valor)


insert_recover = """
                    INSERT INTO recover (id, antenna, checkpoint_id, athlete_num, athlete_time, staff, timestp)
                    SELECT id, antenna, checkpoint_id, athlete_num, athlete_time, staff, timestp FROM athletes_times
                """

delete_from_athletes_times = "DELETE FROM athletes_times"
