from sanic import Sanic
from service.handlers import register_handlers
import os

app = Sanic('tabnewsRSS')
register_handlers(app)

if __name__ == '__main__':
    app.go_fast(
        port=int(os.environ.get('PORT', 8080)),
        workers=os.cpu_count() or 1
    )
