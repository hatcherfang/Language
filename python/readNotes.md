## python dict  
python的实现是采用散列表(hash table)。解决冲突的方式是开放定址法。  
在删除元素时时实行**伪删除(dummy态)**操作，因为这样起到承前启后的作用，为了元素正确查找。  
同时把伪删除的entry 放在缓存池中，有新元素插入的时候，首先会查找缓存池，如果有空余会预先在缓存池中插入新元素。  
Dict是类似C++ STL中map(实现基于红黑树)的一种关联容器。  
而关联容器中一个(键,值)元素对称为一个entry或slot, 记为PyDictObject。  
```
typedef struct {
    Py_ssize_t me_hash; /*cached hash code of me_key*/
    PyObject *me_key;
    PyObject *me_value;
} PyDictEntry;
```
PyDictObject中其实存放的都是PyObject*。**所以dict类型的变量在python中进行变量传递的时候，如果改变dict内的元素值，那么传递前的dict内的元素也会改变，因为这种传递类似引用传递**在PyDictEntry中, me_hash域存储的是me_key的散列值，利用  
一个域来记录这个散列值可以避免每次查询的时候都要重新计算一遍散列值。    
在python中，关联容器是通过PyDictObject对象来实现的。而一个PyDictObject对象实际上是一大堆entry的集合。  
在创建Dict时，会被分配一个entry个数为8的数组，这个值是经过大量实验得出的最佳值。它既不会太浪费内存空间，又能很好满足Python内部大量使用  
PyDictObject的环境的要求，不需要在使用的过程中再次调用malloc申请内存空间。  
当table的装载率大于等于2/3时，table的大小需要变动。规则如下：  
```
mp->ma_used*(mp->ma_used>50000 ? 2 : 4)
```
即entry used态大于50000创建新的table的大小按照2倍来扩充，否则按照4倍来扩充。  
调整table大小也会向size 变小的方向调整。  
当将旧table数据copy到新table的时候，会忽略dummy状态的entry，因为新table中不需要伪删除来查找元素。copy完之后会把旧table空间释放。  
