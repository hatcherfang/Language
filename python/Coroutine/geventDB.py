import sqlalchemy
from sqlalchemy import text
import time
import gevent
import gevent.monkey
gevent.monkey.patch_all(dns=gevent.version_info[0]>=1)


def connect():
    SQL_ENGINE_URL_FORMAT = '{dialect}+{driver}://{user}:{password}@{host}/{dbname}'
    # for mysql database
    DB_HOST = 'host ip'
    DB_DBNAME = 'dbname'
    DB_ACCOUNT = 'root'
    DB_PASSWORD = 'password'
    DB_DRIVER = 'mysqldb'
    DB_DIALECT = 'mysql'
    SQL_ENGINE_URL = SQL_ENGINE_URL_FORMAT.format(
        dialect=DB_DIALECT,
        driver=DB_DRIVER,
        user=DB_ACCOUNT,
        password=DB_PASSWORD,
        host=DB_HOST,
        dbname=DB_DBNAME)
    # print (setting.DB_HOST)
    # set pool size to 5
    # mysql server will close idling database connection after 8 hours,
    # set recycle as 3600s to reconnect every 1 hour
    engine = sqlalchemy.create_engine(SQL_ENGINE_URL,
                                      pool_size=5,
                                      pool_recycle=3600)
    return engine


def execSql(engine, sql):
    engine = connect()
    result = engine.execute(text(sql))
    print result.fetchall()
    time.sleep(1)

if __name__ == '__main__':
    '''set global database connection handle and to prove that using gevent
    will not increase database connection numbers'''
    engine = connect()
    sql = "select * from tablename"
    greenlets = []
    for i in xrange(10):
        greenlet = gevent.spawn(execSql, engine, sql)
        greenlets.append(greenlet)
    gevent.joinall(greenlets)
    time.sleep(10)
