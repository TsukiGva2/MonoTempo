from time import sleep
import setup
from processa_atletas import processa_atletas


while True:
    processa_atletas()
    sleep(setup.TEMPO_ENVIO)
