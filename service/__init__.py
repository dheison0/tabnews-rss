from os import environ
from dotenv import load_dotenv


load_dotenv()

API_BASE = 'https://www.tabnews.com.br/api/v1'
DATABASE_URI = environ.get("DATABASE_URI", "")
DESCRIPTION = "Postagens mais recentes de seus perfis preferidos!"
SITE = 'https://www.tabnews.com.br'
TITLE = "TabNews RSS"
