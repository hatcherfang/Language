import requests
import json
import gevent
# add monkey patch to improve cocurrency
from gevent import monkey
monkey.patch_all()
# global variable
num = 0


def requestsGet(url):
    print("get: %s" % url)
    # gevent.sleep(0)
    data = requests.get(url)
    ret = data.text
    print(url, len(ret), data.encoding)
    if 'timeline' in url:
        print(ret)
        print("\n")
        print(data.content)
        print("\n")
        print(data.json())


def requestsPost(url, body):
    # need to declare the global variable
    global num
    print("post: %r" % url)
    # gevent.sleep(0)
    r = requests.post(url, data=body)
    print(url, len(r.text))
    # print(r.text)
    # print(help(r))
    num = num + 1
    print r.close()
    # print("\n")


def sendGet():
    jobs = []
    jobs.append(gevent.spawn(requestsGet, 'https://www.python.org/'))
    jobs.append(gevent.spawn(requestsGet, 'https://www.yahoo.com/'))
    jobs.append(gevent.spawn(requestsGet, 'https://github.com/'))
    jobs.append(gevent.spawn(requestsGet, 'https://github.com/timeline.json'))
    gevent.joinall(jobs)


def sendPost():
    jobs = []
    body = {'key1': 'value1', 'key2': 'value2'}
    jobs.append(gevent.spawn(requestsPost, 'http://httpbin.org/post', body))
    # # json.dumps to make the body to string or it will be a form
    # jobs.append(gevent.spawn(requestsPost, 'http://httpbin.org/post',
    #                          json.dumps(body)))
    gevent.joinall(jobs)

if __name__ == '__main__':
    jobs = []
    for i in xrange(1000):
        jobs.append(gevent.spawn(sendPost))
    gevent.joinall(jobs)
    print num
