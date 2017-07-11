# Linux system command
- check centos version
>> `cat /etc/redhat-release`
- vim batch replace
>> `:%s/source_pattern/target_pattern/g`  
- check uid gid  
```
  [root@anyone-pc ROOT]# id ddwi  
  uid=1000(ddwi) gid=1000(ddwi) groups=1000(ddwi)  
  [root@anyone-pc ROOT]# id root  
  uid=0(root) gid=0(root) groups=0(root)  
```
- awk: delete all the python process
>> `ps -ef | grep python | awk '{print $2}' | xargs kill -9` 
