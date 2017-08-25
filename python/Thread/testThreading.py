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
import time


def worker(num):
    """
    thread worker function
    :return:
    """
    time.sleep(10)
    print("The num is  %d" % num)
    return
threads = []
for i in range(10):
    t = threading.Thread(target=worker, args=(i,), name="t.%d" % i)
    threads.append(t)
    t.start()
for thread in threads:
    thread.join()
print "end"
