class Test(object):
    def hello(self):
        print("hello, world")

if __name__ == '__main__':
    t = Test()
    t.hello()
    # help(t)
    print dir(t)
    print type(t)
    print type(type(t))
    type(t)().hello()
