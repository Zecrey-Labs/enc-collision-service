#!/bin/bash
wrkBench(){
    echo "---------TESTING URL $i START-----------"
    wrk -t8 -c100 -d10s -s post.lua  --latency http://172.28.218.178:8890/api/v1/crypto/getEncCollision;
    echo "---------TESTING URL $i FINISH-----------"
    echo ""
}
wrkBench >> postoutput.txt

