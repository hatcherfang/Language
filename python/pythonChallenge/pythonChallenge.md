### [Python Challenge](http://www.pythonchallenge.com)  
0. pow(2, 38)  
1.   
```
s = "g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle\
    gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle.\
    sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj."
s = s.lower()
cs = map(lambda s: chr(ord(s) + 2) if s >= 'a' and s < 'y' else 'a'
         if s == 'y' else 'b' if s == 'z' else s, s)
print ''.join(cs)
cs = map(lambda s: chr(ord(s) + 2) if s >= 'a' and s < 'y' else 'a'
         if s == 'y' else 'b' if s == 'z' else s, "map")
print ''.join(cs)
```
2.  
```
import re
fp = open("file_2")
result = []
for line in fp:
    s = re.findall(r'[a-zA-Z]', line)
    if s:
        result.extend(s)
print ''.join(result)
```
3. Reference(https://the-python-challenge-solutions.hackingnote.com/level-3.html)  
```
import re
fp = open("file_3")
result = []
for line in fp:
    s = re.findall(r'[^A-Z]+[A-Z]{3}([a-z])[A-Z]{3}[^A-Z]+', line)
    if s:
        result.extend(s)
s = ''.join(result)
print s
```
