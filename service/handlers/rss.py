import feedparser
from sanic.response import text
from http import HTTPStatus
from aiohttp import ClientSession
from ..database import db

async def get_user_posts(user_name: str):
    session = ClientSession()
    response = await session.get(f"{API_BASE}/contents/{user_name}")


async def rss_feed(_):
    return text("")

