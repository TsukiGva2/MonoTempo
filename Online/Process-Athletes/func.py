import logging
from typing import Never

from setup import DEBOUG

logging.basicConfig(level=logging.INFO)  # Set logging level


def log(data):
    if DEBOUG:
        logging.info(data)


def err(e: Exception) -> Never:
    logging.error(str(e))
    raise e


def suc(data):
    logging.info(f"\033[32;1m{data}\033[0m")


def warn(data):
    if DEBOUG:
        logging.warning(data)
