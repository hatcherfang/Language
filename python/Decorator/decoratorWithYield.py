import requests


def retry(times):
    def decorator(func):
        def wrapper(*args, **kw):
            i = 0
            while i < times:
                yield func(*args, **kw)
                i = i + 1
        return wrapper
    return decorator


@retry(3)
def yieldFunc():
    yield requests.get('http://www.python.org/')

if __name__ == "__main__":
    for i in yieldFunc():
        print i.next().status_code
        print i
