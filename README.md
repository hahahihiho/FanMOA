# FanMOA

## Structure
* Hyperledger-fabric
* Chaincode
* Web

## What you need to prepare for 
* hyperledger-fabric
	* linux(ubuntu)
	* cryptogen, configtxgen
	* docker,docker-compose

* chaincode
	* golang

* web
	* nodejs (express, ejs)

## How to start it
* hyperledger-fabric
    1. generate.sh (make config files)
    2. start.sh (create network)
* chaincode
	* deploy_cc.sh 
* web
	* run enrollAdmin, registerUser
	* run index.js