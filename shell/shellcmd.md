# Linux System Command
- check centos version
```
cat /etc/redhat-release
```
- vim batch replace
```
:%s/source_pattern/target_pattern/g  
```
- check uid gid  
```
[root@anyone-pc ROOT]# id ddwi  
uid=1000(ddwi) gid=1000(ddwi) groups=1000(ddwi)  
[root@anyone-pc ROOT]# id root  
uid=0(root) gid=0(root) groups=0(root)  
```
- awk: delete all the python process
```
ps -ef | grep python | awk '{print $2}' | xargs kill -9 
```
- unzip: unzip zip file and override exists files
```
unzip -o filename.zip 
```
- screen management tool: tmux
```
yum install tmux  
```
Common command:  
1. create session: `tmux new -s s1 -n w1`  
2. create new session: `ctrl + b + c`  
3. show session list(must be in session): `tmux ls`  
4. split window: `ctrl + b + %`, `ctrl + b + "`  
5. minimum/maximum window: `ctrl + b + z`  
6. switch window: `ctrl + b + o`, `ctrl + b + q` + panel number  
7. close panel: `ctrl + b + x`
8. move the current pane left `ctrl + b + {`
9. move the current pane right `ctrl + b + }`
- check file/directory disk size  
```
du -ah file/directory  

run command `man du` to show parameters as below:

...
-a, --all
       write counts for all files, not just directories

-h, --human-readable
       print sizes in human readable format (e.g., 1K 234M 2G)

-s, --summarize
       display only a total for each argument
...
```
- upload/download files sz/rz command  
1. yum install lrzsz  
2. sz means send files to client and rz means receive files from client and these operation happen at server that you install the package 
3. z may be means zmodem a file transfer protocal  
4. usage:  
`sz filename1 filename2` to send server files to client  
`rz` to receive client files, that means upload client files to server  

- cassandra command in linux environment  
1. `cqlsh`  to get into cassandra  
2. `describe keyspaces`  to show tables space  
3. `use keyspacename` to get into keyspace name
4. `describe tables`  to show tables's name  
5. `describe table table_name`  to show table structure  


**Related Tutorials**:   
- [优雅地使用命令行：Tmux 终端复用](http://harttle.com/2015/11/06/tmux-startup.html)  
- [Linux下终端利器tmux](http://kumu-linux.github.io/blog/2013/08/06/tmux/)  
- [Linux基础：利用SSH上传、下载（使用sz与rz命令）](http://skypegnu1.blog.51cto.com/8991766/1538371)  
