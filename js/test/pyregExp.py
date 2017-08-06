import re
# convert regExp to pattern object
s = r'\d{3}\s*\-\s*\d{3,8}'
pattern = re.compile(s)
d = '012 - 23232'
match = pattern.match(d)
if match:
    print match.group()

#r = 'a b  c'.split([\s+])
#print r
