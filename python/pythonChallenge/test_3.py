import re
fp = open("file_3")
result = []
for line in fp:
    s = re.findall(r'[^A-Z]+[A-Z]{3}([a-z])[A-Z]{3}[^A-Z]+', line)
    if s:
        result.extend(s)
s = ''.join(result)
print s
