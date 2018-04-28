import multiprocessing as mp
import time


class workerWithLock():
    'multiprocess with lock'
    def __init__(self, value, interval, lock):
        self.value = value
        self.interval = interval
        self.lock = lock

    def worker(self):
        self.lock.acquire()
        print "workerWithLock"
        time.sleep(self.interval)
        self.value.value += 1
        print self.value.value
        print "end workerWithLock"
        self.lock.release()


class workerWithoutLock():
    'multiprocess without lock'
    def __init__(self, value, interval, lock=None):
        self.value = value
        self.interval = interval

    def worker(self):
        print "workerWithoutLock"
        time.sleep(self.interval)
        self.value.value += 1
        print self.value.value
        print "end workerWithoutLock"


def testWorker(cworker, value, interval, pnum, lock=None):
    print cworker.__doc__ + "\n"
    worker = cworker(value, interval, lock)
    for _ in xrange(pnum):
        p = mp.Process(target=worker.worker)
        p.start()


def multiprocessAttr():
    print("The number of CPU is:" + str(mp.cpu_count()))
    for p in mp.active_children():
        print("child   p.name:" + p.name + "\tp.id" + str(p.pid))
    print "END!!!!!!!!!!!!!!!!!"

if __name__ == "__main__":
    v1 = mp.Value("i", 0)  # "i" means integer type
    interval = 0
    lock = mp.Lock()
    pnum = 10000
    testWorker(workerWithLock, v1, interval, pnum, lock=lock)
    # testWorker(workerWithoutLock, v1, interval, pnum, lock=lock)
