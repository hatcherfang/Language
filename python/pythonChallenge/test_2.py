import re
fp = open("file_2")
result = []
for line in fp:
    s = re.findall(r'[a-zA-Z]', line)
    if s:
        result.extend(s)
print ''.join(result)
