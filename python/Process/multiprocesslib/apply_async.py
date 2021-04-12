from multiprocessing import Pool
import time


def test(p):
    print(p)
    time.sleep(3)


if __name__ == "__main__":
    pool = Pool(processes=2)
    for i in range(500):
        '''
         （1）循环遍历，将500个子进程添加到进程池（相对父进程会阻塞）\n'
         （2）每次执行2个子进程，等一个子进程执行完后，立马启动新的子进程。（相对父进程不阻塞）\n'
        '''
        pool.apply_async(test, args=(i,))   # 维持执行的进程总数为10，当一个进程执行完后启动一个新进程.
    print('test')
    pool.close()
    pool.join()
