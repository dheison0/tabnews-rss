from . import API_BASE
from ..database import db
from aiohttp import ClientSession
from dataclasses import asdict
from http import HTTPStatus
from sanic.request import Request
from sanic.response import json


async def add_user(request: Request):
    username = request.args.get('user')
    if not username:
        return json(
            {'error': 'missing "user" query entry'},
            status=HTTPStatus.BAD_REQUEST
        )
    session = ClientSession()
    response = await session.get(
        f'{API_BASE}/contents/{username}',
        params={'page': 1, 'per_page': 1}
    )
    await session.close()
    del session
    if response.status != HTTPStatus.OK:
        return json({'error': 'user not found'}, status=HTTPStatus.BAD_REQUEST)
    db.add_user(username)
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

