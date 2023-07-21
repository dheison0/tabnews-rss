import asyncio
import rfeed

from . import API_BASE, SITE
from .. import TITLE, DESCRIPTION
from ..database import db
from aiohttp import ClientSession
from datetime import datetime
from http import HTTPStatus
from sanic.response import text
from time import strptime
from typing import Union


async def get_user_posts(user_name: str) -> Union[str, list[dict] | None]:
    session = ClientSession()
    response = await session.get(f"{API_BASE}/contents/{user_name}")
    data = await response.json()
    await session.close()
    return user_name, data


def turn_post_into_feed_item(post: dict) -> rfeed.Item:
    link = f"https://www.tabnews.com.br/{post['owner_username']}/{post['slug']}"
    time = strptime(post['published_at'], '%Y-%m-%dT%H:%M:%S.%fZ')
    publish_date = datetime(
        year=time.tm_year,  month=time.tm_mon,
        day=time.tm_mday,   hour=time.tm_hour,
        minute=time.tm_min, second=time.tm_sec
    )
    return rfeed.Item(
        title=post['title'],
        link=link,
        author=post['owner_username'],
        pubDate=publish_date,
        guid=rfeed.Guid(link)
    )


async def rss_feed(_):
    users = db.get_users()
    tasks = [get_user_posts(u.name) for u in users]
    raw_results = await asyncio.gather(*tasks)
    items = []
    for result in raw_results:
        if result[1] is None:
            db.update_user_status(result[0], 'not found')
            continue
        for p in result[1]:
            if p['title'] == None:
                continue
            items.append(turn_post_into_feed_item(p))
    feed = rfeed.Feed(
        title=TITLE,
        description=DESCRIPTION,
        link=SITE,
        items=items,
        language='pt_BR',
    )
    rss = feed.rss()
    return text(rss, status=HTTPStatus.OK)

