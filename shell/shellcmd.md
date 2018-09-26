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
10. [root@sm ~]# tmux kill-session -t 0  
usage: kill-session [-a] [-t target-session]  
11. tmux at -t session  

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
eg:  
du -h --max-depth=1 directory  
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

- check certificate info  
```
# 查看KEY信息
> openssl rsa -noout -text -in myserver.key

# 查看CSR信息
> openssl req -noout -text -in myserver.csr

# 查看证书信息
> openssl x509 -noout -text -in ca.crt

# 验证证书
# 会提示self signed
> openssl verify selfsign.crt

# 因为myserver.crt 是幅ca.crt发布的，所以会验证成功
> openssl verify -CAfile ca.crt myserver.crt
```
- 判断一个证书是否是自签证书的方法  
1. 查看证书信息中的Issuer(发行者)即可。  

- Nslookup is a program to query Internet domain name servers.   
Usage:  
查看ip地址为ipA的domain name  
`nslookup ipA`  

- 查看linux所有网卡信息  
`ip a`  

- 查询创建文件的原始rpm包：  
rpm -qf /var/directory(directory文件夹名称)  

- 如何对一个卡死的程序进行debug  
1. 查看messages log `vim /var/log/messages`  
2. 查看卡死的进程pid, 根据pid追综进程栈行为  
`ps -ef | grep 进程名`    
`pstack "pid"`  
3. 查看进程正在进行的行为  
`strace -p "pid"`  
4. 根据pid到/proc/pid/fd/目录下查看  
`ll /proc/pid/fd/`  

- [tee命令](http://man.linuxde.net/tee)  
1. tee命令用于将数据重定向到文件，另一方面还可以提供一份重定向数据的副本作为后续命令的stdin。  
简单的说就是把数据重定向到给定文件和屏幕上。  
![tee workflow](http://man.linuxde.net/wp-content/uploads/2013/12/073315SF8.gif)  
存在缓存机制，每1024个字节将输出一次。若从管道接收输入数据，应该是缓冲区满，才将数据转存到指定的文件中。  
若文件内容不到1024个字节，则接收完从标准输入设备读入的数据后，将刷新一次缓冲区，并转存数据到指定文件。  
tee(option)(param)  
option:  
-a：向文件中重定向时使用追加模式；  
-i：忽略中断（interrupt）信号。  
param:  
filename：指定输出重定向的文件。
- [split命令](http://man.linuxde.net/split)  
1. split命令可以将一个大文件分割成很多个小文件，有时需要将文件分割成更小的片段，比如为提高可读性，生成日志等。  
split option  

- [awk](http://www.zsythink.net/archives/1357)   
1. awk split string  
eg:  
file.txt  
```
abc#123#def
abc#123#def
abc#123#def
```
Run command as below:  
`cat file.txt | awk -F# '{print $1, $2, $3}'`  
OR:  
`awk -F# '{print $1, $2, $3}' file.txt`  
Output:  
```
abc 123 def
abc 123 def
abc 123 def
```
- [sort](https://blog.csdn.net/monkeyduck/article/details/10097829)  
1. sort的工作原理  
sort将文件的每一行作为一个单位，相互比较，比较原则是从首字符向后，依次按ASCII码值进行比较，最后将他们按升序输出。  
eg:  
```
[rocrocket@rocrocket programming]$ cat seq.txt
banana
apple
pear
orange
[rocrocket@rocrocket programming]$ sort seq.txt
apple
banana
orange
pear
```
2. sort的-u选项  
它的作用很简单，就是在输出行中去除重复行。  
```
[rocrocket@rocrocket programming]$ cat seq.txt
banana
apple
pear
orange
pear
[rocrocket@rocrocket programming]$ sort seq.txt
apple
banana
orange
pear
pear
[rocrocket@rocrocket programming]$ sort -u seq.txt
apple
banana
orange
pear

pear由于重复被-u选项无情的删除了。
```
3. sort的-r选项  
sort默认的排序方式是升序，如果想改成降序，就加个-r就搞定了。  
```
[rocrocket@rocrocket programming]$ cat number.txt
1
3
5
2
4
[rocrocket@rocrocket programming]$ sort number.txt
1
2
3
4
5
[rocrocket@rocrocket programming]$ sort -r number.txt
5
4
3
2
1
```
4. sort的-o选项  
由于sort默认是把结果输出到标准输出，所以需要用重定向才能将结果写入文件，形如sort filename > newfile。  
但是，如果你想把排序结果输出到原文件中，用重定向可就不行了。  
```
[rocrocket@rocrocket programming]$ sort -r number.txt > number.txt
[rocrocket@rocrocket programming]$ cat number.txt
[rocrocket@rocrocket programming]$
```
看，竟然将number清空了。  
就在这个时候，-o选项出现了，它成功的解决了这个问题，让你放心的将结果写入原文件。这或许也是-o比重定向的唯一优势所在。  
```
[rocrocket@rocrocket programming]$ cat number.txt
1
3
5
2
4
[rocrocket@rocrocket programming]$ sort -r number.txt -o number.txt
[rocrocket@rocrocket programming]$ cat number.txt
5
4
3
2
1
```
- ssh登陆连接配置port和ssh key  
1. configure port  
open file `vim /etc/ssh/ssh_config`   
add record:  
```
Host xx*
    Port xxx
```
2. copy `id_rsa` to path ~/.ssh/  
3. `chmod 400 ~/.ssh/id_rsa`  
**Related Tutorials**:   
- [优雅地使用命令行：Tmux 终端复用](http://harttle.com/2015/11/06/tmux-startup.html)  
- [Linux下终端利器tmux](http://kumu-linux.github.io/blog/2013/08/06/tmux/)  
- [Linux基础：利用SSH上传、下载（使用sz与rz命令）](http://skypegnu1.blog.51cto.com/8991766/1538371)  
- [常用Openssl命令](http://www.cnblogs.com/E7868A/archive/2012/11/16/2772240.html)  
- [Linux通过PID查看进程完整信息](http://blog.csdn.net/great_smile/article/details/50114133)  
