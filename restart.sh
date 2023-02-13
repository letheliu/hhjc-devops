#!/bin/bash

# restart zihao

# kill zihao process
ps -ef | grep ./zihao | grep -v 'grep' | grep -v 'restart.sh' | awk '{print $2}' | xargs kill -9
cd /zihao/master
chmod u+x zihao
# start zihao

./zihao > zihao.log &



