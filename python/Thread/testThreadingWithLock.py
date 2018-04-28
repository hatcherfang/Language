#!/usr/bin/python
# coding=utf-8
import threading
import time
num = 0
mutex = threading.Lock()


class MyThreadWithLock(threading.Thread):
    def run(self):
        global num, mutex
        time.sleep(1)
        if mutex.acquire(1):
            num = num+1
            msg = self.name + ' set num to ' + str(num)
            print msg
            mutex.release()


class MyThreadWithNoLock(threading.Thread):
    def run(self):
        global num
        time.sleep(1)
        num = num+1
        msg = self.name + ' set num to ' + str(num)
        print msg


def testMyThreadWithLock():
    for i in range(100):
        t = MyThreadWithLock()
        t.start()


def testMyThreadWithNoLock():
    for i in range(100):
        t = MyThreadWithNoLock()
        t.start()
if __name__ == '__main__':
    testMyThreadWithLock()
    # testMyThreadWithNoLock()
