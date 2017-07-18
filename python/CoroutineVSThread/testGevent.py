import gevent
from gevent import monkey
monkey.patch_all(dns=gevent.version_info[0]>=1)
import time
import random
globals_num = 0


def Func():
    global globals_num
    time.sleep(random.randint(0, 2))
    globals_num += 1
    #print(globals_num)

gevents = []
for i in range(1000):
    t = gevent.spawn(Func)
    gevents.append(t)
gevent.joinall(gevents)
print(globals_num)
