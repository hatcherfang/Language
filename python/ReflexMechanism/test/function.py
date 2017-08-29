'''
usage:
    python function.py
'''
import reflex
class factory1(object):
    def create_factory(self):
        print "concrete product1"


class factory2(object):
    def create_factory(self):
        print "concrete product2"

if __name__ == "__main__":
    factory_list = reflex.getFactory().getAllFactory()
    for factory in factory_list:
        factory().create_factory()
