#!/bin/bash
set -e
RESET=false
if [ $RESET = true ]; then
    ./teardown.sh
    ./start.sh
    ./deploy_cc.sh
fi
# 변수설정
# CC_SRC_PATH=github.com/paper-contract/
# docker base_path : /opt/gopath/src/
CC_SRC_PATH=chaincode/fanmoa/
CHANNEL_NAME=mychannel
CC_NAME=fanmoa
VERSION=0.9
CC_RUNTIME_LANGUAGE=go

# maybe in the future
# makeCommand(){
#     cmd=`{"Args":["$1",`
#     return $cmd
# }

# test
if [ $RESET = false ]; then
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerEvent","event2","evvent222!!!","user2","user2","20220110","20220301","1000","100000"]}'

# reset
else
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerUser","user1"]}'
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerUser","user2"]}'
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerUser","user3"]}'
    sleep 3
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["getUser","user1"]}'
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["getUser","user2"]}'
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["getUser","user3"]}'
    sleep 3


    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerEvent","event1","evvent!","user1","user2","20211010","20220101","100","1000"]}'
    sleep 3
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["registerEvent","event2","evvent222!!!","user2","user2","20220110","20220301","1000","100000"]}'
    sleep 3

    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["putMoney","user3","event1"]}'
    sleep 3
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["completeEvent","event1"]}'
    sleep 3

    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["getUserHistory","user1"]}'
    docker exec cli peer chaincode invoke -n $CC_NAME -C $CHANNEL_NAME -c '{"Args":["getAllEvents"]}'
fi

