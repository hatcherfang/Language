#!/usr/bin/env python
# -*- coding: utf-8 -*-
import requests
import gevent
from gevent import monkey, pool
monkey.patch_all()


def get(url, index):
    res = requests.get(url)
    print(res.status_code, index)
    return res.status_code == 200


def test():
    url = "https://wwww.baidu.com"
    geventjobs = []
    p = pool.Pool(20)
    chunk_count = 100
    index = 0
    while index < chunk_count:
        ret = p.spawn(get, url, index)
        index += 1
        geventjobs.append(ret)
    gevent.joinall(geventjobs)
    for job in geventjobs:
        if job.value is False:
            print('request error')


if __name__ == "__main__":
    test()
