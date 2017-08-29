# method1 __import__ package string name
import json
class getFactory(object):
    def __init__(self):
        with open("./file.json") as fp:
            self.data = fp.read()

    def getAllFactory(self):
        self.data = json.loads(self.data)
        print self.data
        factory_list = []
        if self.data:
            cc = __import__(self.data['moduleName'])
            className = self.data['className']
            for c in className:
                if hasattr(cc, c):
                    factory_list.append(getattr(cc, c))
        return factory_list
# method2 getattr string function name
# class factory(object):
#     def create_factory(self):
#         print "concrete product"
# cc = factory()
# if hasattr(cc, "create_factory"):
#     c = getattr(cc, "create_factory")
#     print c()



