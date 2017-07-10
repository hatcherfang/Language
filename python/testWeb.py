import web

urls = (
    '/(.*)', 'hello'
)

class hello:
    def GET(self, name):
        i = web.input(times=1)
        if not name: name = 'world'
        for c in xrange(int(i.times)):
            return 'Hello,', name+'!'

if __name__ == "__main__":
    app = web.application(urls, globals())
    application = app.wsgifunc()
    app.run()
