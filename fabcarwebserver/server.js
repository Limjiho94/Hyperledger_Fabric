// ExpressJS Setup
const path = require('path');
const express = require('express');
const app = express();
var bodyParser = require('body-parser');

// Hyperledger Bridge Setup
const { Wallets, Gateway } = require('fabric-network');
const fs = require('fs');

// load the network configuration
const ccpPath = path.resolve('/home/apstudent', 'fabric-samples', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
const ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));


// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);

// use static file
app.use(express.static(path.join(__dirname)));

// configure app to use body-parser
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

// main page routing
app.get('/', function(req, res){
    res.sendFile(__dirname + '/viewpage/index.html');
})
app.get('/page/queryTotalInfo', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryTotalInfo.html');
})
app.get('/page/queryAll', function(req, res){    
    res.sendFile(__dirname + '/viewpage/queryAll.html');
})
app.get('/page/setManufactureInfo', function(req, res){    
    res.sendFile(__dirname + '/viewpage/setManufactureInfo.html');
})
app.get('/page/setGovernmentInfo', function(req, res){    
    res.sendFile(__dirname + '/viewpage/setGovernmentInfo.html');
})
app.get('/page/setRepairInfo', function(req, res){    
    res.sendFile(__dirname + '/viewpage/setRepairInfo.html');
})
app.get('/page/setInsuranceInfo', function(req, res){    
    res.sendFile(__dirname + '/viewpage/setInsuranceInfo.html');
})
app.get('/page/changeCarOwner', function(req, res){    
    res.sendFile(__dirname + '/viewpage/changeCarOwner.html');
})

// api routing
app.get('/api/queryAll', async function(req, res){
    const result = await callChainCode('queryAll')        
    res.json(JSON.parse(result))    
})

app.post('/api/queryTotalInfo', async function(req, res){
    const carno=req.body.carno
    const result = await callChainCode('queryTotalInfo',carno)    
    res.json(JSON.parse(result))
})

app.post('/api/setManufactureInfo', async function(req, res){
    const carno = req.body.carno
    const carPIN = req.body.carPIN
    const carCCvol = req.body.carCCvol
    const carmodel = req.body.carmodel
    const carprice = req.body.carprice
    const carcolor = req.body.carcolor

    var args = [carno,carPIN,carCCvol,carmodel,carprice,carcolor]   
    //console.log(`result:${args}`); 
    await callChainCode('setManufactureInfo',args)    
    res.status(200).json({result: "success"})
})

app.post('/api/setGovernmentInfo', async function(req, res){
    const carno = req.body.carno
    const carowner = req.body.carowner
    const carownerAdd = req.body.carownerAdd
    const carownerNumber = req.body.carownerNumber

    var args = [carno,carowner,carownerAdd,carownerNumber]
    
    
    await callChainCode('setGovernmentInfo',args)    
    res.status(200).json({result: "success"})
})

app.post('/api/setRepairInfo', async function(req, res){
    const carno = req.body.carno
    const repairhistory = req.body.repairhistory
    const repairdate = req.body.repairdate
    const repairplace = req.body.repairplace
    const mechanicname = req.body.mechanicname

    var args = [carno,repairhistory,repairdate,repairplace,mechanicname]
    
    
    await callChainCode('setRepairInfo',args)    
    res.status(200).json({result: "success"})
})

app.post('/api/setInsuranceInfo', async function(req, res){
    const carno = req.body.carno
    const subnumber = req.body.subnumber
    const productInfo = req.body.productInfo
    const subscriber = req.body.subscriber
    const managerName = req.body.managerName
    const companyName = req.body.companyName

    var args = [carno,subnumber,productInfo,subscriber,managerName,companyName]
    
    
    await callChainCode('setInsuranceInfo',args)    
    res.status(200).json({result: "success"})
})


app.post('/api/changeCarOwner', async function(req, res){
    const carno = req.body.carno
    const carowner = req.body.carowner

    var args = [carno,carowner]
    await callChainCode('changeCarOwner',args)    
    res.status(200).json({result: "success"})
})

async function callChainCode(fnName, args){
    
    // Create a new file system based wallet for managing identities.
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = await Wallets.newFileSystemWallet(walletPath);
    var result;
    console.log(`Wallet path: ${walletPath}`);
    

    // Check to see if we've already enrolled the user.
    const identity = await wallet.get('appUser');
    if (!identity) {
        console.log('An identity for the user "appUser" does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }
    
    // Create a new gateway for connecting to our peer node.
    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });
    
    // Get the network (channel) our contract is deployed to.
    const network = await gateway.getNetwork('mychannel');

    // Get the contract from the network.
    const contract = network.getContract('UsedCar');
    
    // Evaluate the specified transaction.    
    if(fnName == 'queryAll')
        result = await contract.evaluateTransaction(fnName);    
    else if(fnName == 'queryTotalInfo')
        result = await contract.evaluateTransaction(fnName,args);
    else if(fnName == 'setManufactureInfo')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4],args[5])
    else if(fnName == 'setGovernmentInfo')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3])
    else if(fnName == 'setRepairInfo')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4])
    else if(fnName == 'setInsuranceInfo')
        result = await contract.submitTransaction(fnName,args[0],args[1],args[2],args[3],args[4],args[5])
    else if(fnName == 'changeCarOwner')
        result = await contract.submitTransaction(fnName,args[0],args[1])
    else
        result = 'This function(' + fnName +') does not exist !'        
        
    console.log(`Transaction has been evaluated, result is: ${result.toString()}`);

    // Disconnect from the gateway.
    await gateway.disconnect();
    
    return result;
}
