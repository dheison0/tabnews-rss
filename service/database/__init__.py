from .local import LocalDatabase
from .. import DATABASE_URI


db = LocalDatabase(DATABASE_URI)

