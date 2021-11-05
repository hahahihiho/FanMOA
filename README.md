# FanMOA

## Structure
* Linux(ubuntu)
    * Hyperledger-fabric
    * Chaincode(golang)
    * Web(nodejs :express, ejs)

## What we need to install
* install fabric v1.4
	* prerequisites
		- cURL
		- docker&docker-compose
		- GO, GOPATH
	- ref : https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html	
	- ref : https://hyperledger-fabric.readthedocs.io/en/release-1.4/install.html
	

## How to run
* install
	1. install.sh (install hyperledger)(+ docker, golang, nodejs)
	2. setting.sh (symbolic link setting)
* hyperledger-fabric
    1. generate.sh (make config files)
    2. start.sh (create network)
* chaincode
	* deploy_cc.sh 
* web
	* `npm install`
	* ```
	node enrollAdmin
	node registerUser
	```
	* `node index.js`