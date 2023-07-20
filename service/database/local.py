import sqlite3
from .model import DatabaseConnector, User

class LocalDatabase(DatabaseConnector):
    connection: sqlite3.Connection

    def __init__(self, uri: str = ""):
        self.connection = sqlite3.connect(uri)
        self.create_tables()

    def _execute(self, *args, **kwargs):
        cursor = self.connection.cursor()
        cursor.execute(*args, **kwargs)
        cursor.close()
        self.connection.commit()

    def create_tables(self):
        """Create all tables needed to start working"""
        self._execute("CREATE TABLE IF NOT EXISTS users(name TEXT UNIQUE, status TEXT);")

    def get_users(self) -> list[User]:
        cursor = self.connection.cursor()
        cursor.execute("SELECT * FROM users;")
        query_result = cursor.fetchall()
        cursor.close()
        users = list(map(lambda u: User(*u), query_result))
        return users

    def add_user(self, user: str):
        self._execute("INSERT INTO users VALUES (?, ?);", [user, 'Ok'])

    def remove_user(self, user: str):
        self._execute("DELETE FROM users WHERE name=?;", [user])

    def update_user_status(self, user: str, status: str):
        self._execute("UPDATE users SET status=? WHERE name=?", [status, user])


