from dataclasses import dataclass
from typing import Iterable


@dataclass
class User:
    name: str
    status: str = "Ok"


class DatabaseConnector:
    def __init__(self, uri: str):
        """Constructor method that receives an URI to database"""
        pass

    def get_users() -> Iterable[User]:
        """Gets the list of users that you want to receive posts from"""
        pass

    def add_user(user: str):
        """Add a tabnews user to update list"""
        pass

    def remove_user(user: str):
        """Remove user from update list"""
        pass

