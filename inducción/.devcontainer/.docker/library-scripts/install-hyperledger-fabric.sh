#!/bin/bash

CHANNEL_NAME=${1:-"mychannel"}
CCP_FILE_PATH=${2:-"../asset-transfer-basic/chaincode-go"}

CURRENT_DIR=$PWD
DIR=$GOPATH/src/github.com/devcontainer
mkdir -p $DIR
cd $DIR

curl -sSL https://bit.ly/2ysbOFE | bash -s
