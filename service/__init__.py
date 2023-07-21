from os import environ
from dotenv import load_dotenv


load_dotenv()

DATABASE_URI = environ.get("DATABASE_URI", "")
DESCRIPTION = "Postagens mais recentes de seus perfis preferidos!"
TITLE = "TabNews RSS"
