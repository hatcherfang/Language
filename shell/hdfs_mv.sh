#!/bin/bash
srcPath=/logs/tmp
destPath=/logs/fhq
hdfs dfs -ls $srcPath | awk '{print $5, $8}' | while read line
do
    size=`echo $line | cut -d ' ' -f 1`
    filepath=`echo $line | cut -d ' ' -f 1 --complement`
    if [[ $size = 0 ]];
    then
        hdfs dfs -mv $filepath $destPath
    fi
done
