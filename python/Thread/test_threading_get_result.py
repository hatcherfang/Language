'''
tips:
    when we use threading package, we should pay attention that we should not
    using thread.join() after thread.start() for every thread or the thread
    concurrency will not work.

    we should use for cycle to start the thread and then use for cycle to join
    thread just like below.

    we use join() to make parent thread wait until child thread run completed
    and then run the remain code.
    if we don't use join(), the parent thread will not wait and continue run the
    remain code, after the child thread run completed the parent thread will
    stop.

    reference:
        http://blog.csdn.net/zhangzheng0413/article/details/41728869
'''
import threading
from threading import Thread
import time


def worker(num, num2):
    """
    thread worker function
    :return:
    """
    time.sleep(1)
    print("The num is  %d, num2:%d" % (num, num2))
    return num


class CustomThread(Thread):
    def __init__(self, func, *args):
        Thread.__init__(self)
        self.func = func
        self.args = args

    def run(self):
        self.result = self.func(*self.args)

    def get_result(self):
        return self.result



threads = []
for i in range(10):
    t = CustomThread(worker, i, i+1)
    threads.append(t)
    t.start()
for thread in threads:
    thread.join()
    print("result:%r" % thread.get_result())
print "end"
