const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '../config' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

async function submitTx(fn,...args){
    const {success,contract,gateway} = get_contract()
    let result;
    let isSucceed
    const submitFunctions = ["registerUser","registerProject","joinProject","completeProject","recordScore"]
    const evaluateTransactions = ["getUserInfo"]
    if (success){
        if (fn in submitFunctions){
            result = await contract.submitTransaction(fn, ...args);
        }else if (fn in evaluateTransactions){
            result = await contract.evaluateTransaction(fn, ...args);
        }else {
            console.error(`Undefined function : ${fn}`)
            isSucceed = false; result = null;
        }
        console.log('Transaction has been submitted');
        isSucceed = true; result = result;
    }else {
        isSucceed = false; result = null;
    }
    // gateway.disconnect();
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
    const network = await gateway.getNetwork('mychannel');
    const contract = network.getContract('teamate');
    return {"success":true,"result":contract,"gateway":gateway}
}

exports = {get_contract,submitTx};