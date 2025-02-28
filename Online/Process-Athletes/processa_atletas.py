# from config import MYTEMPO_MYSQL_CONFIG
from database.DatabaseController import DatabaseOperation
from func import err, log, suc
from querys import (
    chegada_por_percurso,
    delete_from_athletes_times,
    insert_recover,
    largada_por_percurso,
    percursosSQL,
)


def processa_atletas() -> None:
    with DatabaseOperation() as db:
        try:
            log("Buscando percursos...\n")
            db.execute(percursosSQL, return_as_object=True)

            # pega os percursos da prova
            percursos = db.results

            if db.info["status"] != "success":
                err(ValueError("Ocorreu um erro ao buscar percursos\n"))

            suc(f"Percursos encontrados: {[i.id for i in percursos]}")

            for percurso in percursos:

                # pega os atletas da largada e insere na tabela de resultados
                log("Processando atletas da largada...\n")

                db.execute(largada_por_percurso(percurso.id))
                if db.info["status"] != "success":
                    err(
                        ValueError(
                            f"Erro ao processar atletas da largada (Percurso {percurso.id}) {db.info}"
                        )
                    )

                suc("Atletas da largada processados com sucesso!\n")

                # pega os atletas da chegada e insere na tabela de resultados

                log("Processando atletas da chegada...\n")

                db.execute(chegada_por_percurso(percurso.id))

                if db.info["status"] != "success":
                    err(
                        ValueError(
                            f"Erro ao processar atletas da chegada (Percurso {percurso.id}) {db.info}"
                        )
                    )

                suc("Atletas da largada processados com sucesso!...\n")

            # insere na tabela recover os dados da athletes_times
            db.execute(insert_recover)

            if db.info["status"] != "success":
                err(ValueError("Erro ao salvar dados em 'recover'"))

            # deleta tudo da tabela athletes_times
            db.execute(delete_from_athletes_times)

            if db.info["status"] != "success":
                err(ValueError("Erro ao remover dados de 'athletes_times'"))

        except Exception as e:
            err(e)
