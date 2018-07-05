class Test(object):
    def __init__(self, config):
        self.config = config

    def __getitem__(self, key):
        return self.config[key]


if __name__ == '__main__':
    d = {1: 3, 2: 4}
    t = Test(d)
    print t[1]
