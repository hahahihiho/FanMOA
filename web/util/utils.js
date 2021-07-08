const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '../config' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

const CHANNEL_NAME = "mychannel"
const CONTRACT_NAME = "fanmoa"

async function submitTx(fn,...args){
    const {success,contract,gateway} = await get_contract()
    let result;
    let isSucceed
    const submitFunctions = ["registerUser","registerEvent","putMoney","completeEvent","refundAll"]
    const evaluateTransactions = ["getUser","getEvent","getUserHistory","getAllEvents"]
    if (success){
        try {
            if (submitFunctions.includes(fn)){
                result = await contract.submitTransaction(fn, ...args);
                isSucceed = true;
                console.log('Submit Transaction has been submitted');
            }else if (evaluateTransactions.includes(fn)){
                result = await contract.evaluateTransaction(fn, ...args);
                isSucceed = true;
                console.log('Evaluate Transaction has been submitted');
                result = JSON.parse(result.toString())
            } else {
                console.error(`Undefined function : ${fn}`)
                isSucceed = false; result = null;
            }
        } catch (err){
            console.error(`Could not send Tx : ${err}`);
            isSucceed = false; result = null;
        }
    }else {
        isSucceed = false; result = null;
    }
    gateway.disconnect();
    return {"success":isSucceed,"result":result}
}
// return {success:bool, result:contract, gateway:gateway}
const get_contract = async () => {
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);
    const user = "user1"
    const userExists = await wallet.exists(user);
    if (!userExists) {
        console.log(`An identity for the user "${user}" does not exist in the wallet`);
        console.log('Run the registerUser.js application before retrying');
        return {"success":false,"result":null,"gateway":null};
    }
    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: user, discovery: { enabled: false } });
    const network = await gateway.getNetwork(CHANNEL_NAME);
    const contract = network.getContract(CONTRACT_NAME);
    return {"success":true,"contract":contract,"gateway":gateway}
}

const makeEntry = (success,result,error) => {
    return {success:success, result:result, error: error};
}

exports = module.exports = {submitTx,makeEntry};