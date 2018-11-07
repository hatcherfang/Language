def test():
    print "hello"
    raise Exception("fang")

if __name__ == "__main__":
    try:
        test()
    except Exception as e:
        print e
