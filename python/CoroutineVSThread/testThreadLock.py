import threading
import time
import random
globals_num = 0
lock = threading.RLock()


def Func():
    #lock.acquire()
    global globals_num
    time.sleep(random.randint(0, 2))
    globals_num += 1
    #print(globals_num)
    #lock.release()

threadList = []
for i in range(1000):
    threadList.append(threading.Thread(target=Func))
for th in threadList:
    th.start()
for th in threadList:
    th.join()
print(globals_num)
