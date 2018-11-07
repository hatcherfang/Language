def decorator2(param):
    def decorator1(func):
        def wrapper(*args, **kwargs):
            print(param)
            raise Exception("exception happend")
            return func(*args, **kwargs)
        return wrapper
    return decorator1


def decorator1(func):
    def wrapper(*args, **kwargs):
        return func(*args, **kwargs)
    return wrapper

@decorator2("fanghaiqun")
@decorator1
def hello(name):
    print("hello %s" % name)


if __name__ == '__main__':
    try:
        hello("hatcher")
    except Exception as e:
        print(e)
