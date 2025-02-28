from os import getenv

MYTEMPO_MYSQL_CONFIG = {
    "port": getenv("MYSQL_PORT"),
    "host": getenv("MYSQL_HOST"),
    "user": getenv("MYSQL_USER"),
    "password": getenv("MYSQL_PASS"),
    "database": getenv("MYSQL_DB"),
}
