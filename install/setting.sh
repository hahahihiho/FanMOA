#!/bin/bash

# make symbolic link
BIN=/usr/local/bin/
PWD=${pwd}
FABRIC_BIN=${PWD}/fabric-samples/bin/
NODEJS_BIN=${PWD}/node-v16.13.0-linux-x64/bin/

function makeSymbolicLink() {
  sudo ln -s ${1}${2} ${BIN}${2}
}

function fabricSymbolic() {
    makeSymbolicLink ${FABRIC_BIN} ${1}
}

function nodeSymboic() {
    makeSymbolicLink ${NODEJS_BIN} ${1}
}

fabricSymbolic cryptogen
fabricSymbolic configtxgen

nodeSymboic node
nodeSymboic npm
nodeSymboic npx

# this is for node_modules
sudo apt install build-essential