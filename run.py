from os import environ, cpu_count
from sanic import Sanic
from service.handlers.register import add_handlers


app = Sanic('tabnewsRSS')
add_handlers(app)

if __name__ == '__main__':
    app.go_fast(
        port=int(environ.get('PORT', 8080)),
        workers=cpu_count() or 1
    )
