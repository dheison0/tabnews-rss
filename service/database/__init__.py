from .. import DATABASE_URI
from .local import LocalDatabase

db = LocalDatabase(DATABASE_URI)

