#referance:
#!/bin/bash
s="hatcher,fang,abc"
one=`echo $s | cut -d ',' -f 1`
other=`echo $s | cut -d ',' -f 1 --complement`
echo $one
echo $other 
