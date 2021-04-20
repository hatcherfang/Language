import time
import functools


def decorator2(param):
    def decorator1(func):
        print('func name:%r' % func.__name__)
        def wrapper(*args, **kwargs):
            print(param)
            raise Exception("exception happend")
            return func(*args, **kwargs)
        return wrapper
    return decorator1


def decorator3(func):
    ts = time.time()
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print("args:%r, kwargs:%r", args, kwargs)
        user = "fang"
        # kwargs['user'] = user
        return func(*args, **kwargs)
    print(func.__name__)
    print(time.time()-ts)
    return wrapper


def decorator1(func):
    print('decorator1')
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper


@decorator2("fanghaiqun")
@decorator1
def hello(name):
    print("hello %s" % name)


@decorator3
def test_param(project, name, **kwargs):
    print(kwargs)
    print("project:%r, name:%r", project, name)


if __name__ == '__main__':
    try:
        hello("hatcher")
    except Exception as e:
        print(e)
    test_param("proj", "name")
