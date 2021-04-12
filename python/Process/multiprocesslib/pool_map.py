#!/usr/bin/env python3
# -*- coding:utf-8 -*-
from multiprocessing.pool import Pool


def f(x):
    return x * x


def g(x, y):
    return x**y


def main():
    with Pool(4) as pool:
        result = pool.map(f, [1, 2, 3, 4, 5])
        print(type(result))
        print(result)

    with Pool(4) as pool:
        result = pool.starmap(g, [(1, 3), (2, 4), (3, 5), (4, 6), (5, 7)])
        print(type(result))
        print(result)


if __name__ == '__main__':
    main()
