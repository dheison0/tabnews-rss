from ..database import db
from ..utils import tabnews_user_exists
from dataclasses import asdict
from http import HTTPStatus
from sanic.request import Request
from sanic.response import json
from sqlite3 import IntegrityError


async def add_user(request: Request):
    username = request.args.get('user')
    if not username:
        return json(
            {'error': 'missing "user" query entry'},
            status=HTTPStatus.BAD_REQUEST
        )
    if not await tabnews_user_exists(username):
        return json({'error': 'user not found'}, status=HTTPStatus.BAD_REQUEST)
    try:
        db.add_user(username)
    except IntegrityError:
        return json({'error': 'user already added'}, status=HTTPStatus.BAD_REQUEST)
    return json({'message': 'ok'})


async def get_users(request: Request):
    users = [asdict(u) for u in db.get_users()]
    return json({'message': 'ok', 'users': users})


async def remove_user(request: Request):
    user = request.args.get('user')
    if not user:
        return json(
            {'error': 'query "user" is missing'},
            status=HTTPStatus.BAD_REQUEST
        )
    db.remove_user(user)
    return json({'message': 'ok'})

