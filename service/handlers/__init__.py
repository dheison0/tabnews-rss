from sanic import Sanic

API_BASE = 'https://www.tabnews.com.br/api/v1'

def register_handlers(server: Sanic):
    from . import api
    server.add_route(api.add_user, '/', ['POST'])
    server.add_route(api.get_users, '/', ['GET'])

