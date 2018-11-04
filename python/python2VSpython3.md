## FAQ  
1. python3 not exist long  
```
In Python3, if an integer is within the range of int32 or int64 (depending on how you compiled the Python binary), it is the same as int in Python2. Once it exceeds the range of machine-native integers, it will be automatically converted to what is equal to long in Python2. The process is completely transparent, thus it is safe to just use int(numImages) in Python3.
```
2. python3 not exist xrange  
```
python3 only have `range` and it return an range object to improve performance   
```
3. python3 vs python2 除法  
```
在Python 2.6中，’/’执行传统除法，如果操作数都是整数的话，执行截断的整数除法（即对于结果只保留整数部分），否则，执行浮点除法（保留余数）；’//’执行Floor除法，与Python3.0一样，对于整数执行截断除法，浮点数执行浮点除法。
在Python 3.0中，’/’总是执行真除法，不管操作数的类型，都会返回包含任何余数的浮点结果；’//’执行Floor除法，截除掉余数并且针对整数操作数返回一个整数，如果有任何一个操作数是浮点数，则返回一个浮点数。
```
4. Python3.5中：iteritems变为items  
