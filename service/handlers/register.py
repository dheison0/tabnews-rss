from . import api, rss
from sanic import Sanic


def add_handlers(server: Sanic):
    server.add_route(api.add_user, '/', ['POST'])
    server.add_route(api.get_users, '/', ['GET'])
    server.add_route(rss.rss_feed, '/rss', ['GET'])

