from sanic import Sanic

API_BASE = 'https://www.tabnews.com.br/api/v1'
SITE = 'https://www.tabnews.com.br'

def register_handlers(server: Sanic):
    from . import api, rss
    server.add_route(api.add_user, '/', ['POST'])
    server.add_route(api.get_users, '/', ['GET'])
    server.add_route(rss.rss_feed, '/rss', ['GET'])

