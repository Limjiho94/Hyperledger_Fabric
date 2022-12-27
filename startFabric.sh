#!/bin/bash

set -e

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1


CC_SRC_PATH="${PWD}/chaincode/UsedCar/" # change PATH to my CC PATH
CC_SRC_LANGUAGE="go"

# launch network; create channel and join peer to channel
pushd $HOME/fabric-samples/test-network
./network.sh down
./network.sh up createChannel -ca -s couchdb
./network.sh deployCC -ccn UsedCar -cci initLedger -ccv 1 -ccl ${CC_SRC_LANGUAGE} -ccp ${CC_SRC_PATH}
popd

