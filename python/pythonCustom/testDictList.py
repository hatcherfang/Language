def testDict(dic):
    dic1 = dic
    dic1["a"] = "change"

def testList(li):
    li1 = li
    li1[0] = "change"

if __name__ == '__main__':
    dic = {"a": "haiqun", "b": "fang"}
    print "before:%r" % dic
    testDict(dic)
    print "after:%r" % dic
    li = [1, 2, 3, 4]
    print "before:%r" % li
    testList(li)
    print "after:%r" % li
    # conclusion: dict and list is passed by reference in function parameter
    print "conclusion: dict and list is passed by reference in function parameter"
