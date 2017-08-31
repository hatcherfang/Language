### [Python Challenge](http://www.pythonchallenge.com)  
1. pow(2, 38)  
2. 
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
