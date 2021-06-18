const BSHPerifV1 = artifacts.require("BSHPeripheryV1");
const BSHPerifV2 = artifacts.require("BSHPeripheryV2");
const BSHCoreV1 = artifacts.require("BSHCoreV1");
const BSHCoreV2 = artifacts.require("BSHCoreV2");
const { assert } = require('chai');
const truffleAssert = require('truffle-assertions');
const { deployProxy, upgradeProxy } = require('@openzeppelin/truffle-upgrades');

contract('BSHCore Unit Tests - After Upgrading Contract', (accounts) => {
    let bsh_coreV1, bsh_coreV2;
    var _native = 'PARA';                   var _fee = 10;
    var _uri = 'https://github.com/icon-project/btp'

    before(async () => {
        bsh_coreV1 = await deployProxy(BSHCoreV1, [_uri, _native, _fee]);
        bsh_coreV2 = await upgradeProxy(bsh_coreV1.address, BSHCoreV2);
    });

    it('Re-initialize BSHCore Contract - Failure', async () => {
        await truffleAssert.reverts(
            bsh_coreV2.initialize(_uri, _native, _fee),
            "Initializable: contract is already initialized"
        );
    });

    it(`Scenario 1: Should allow contract's owner to register a new coin`, async () => {
        var _name = "ICON";
        await bsh_coreV2.register(_name);
        output = await bsh_coreV2.coinNames();
        assert(
            output[0] === _native && output[1] === 'ICON'
        );
    });
    
    it('Scenario 2: Should revert when an arbitrary client tries to register a new coin', async () => {   
        var _name = "TRON";
        await truffleAssert.reverts(
            bsh_coreV2.register(_name, {from: accounts[1]}),
            "Unauthorized"
        );
    }); 

    it('Scenario 3: Should revert when contract owner registers an existed coin', async () => {
        var _name = "ICON";
        await truffleAssert.reverts(
            bsh_coreV2.register(_name),
            "ExistToken"
        );
    }); 

    it('Scenario 4: Should allow contract owner to update BSHPeriphery contract', async () => {
        await bsh_coreV2.updateBSHPeriphery(accounts[2]);
    });

    it('Scenario 5: Should revert when arbitrary client updates BSHPeriphery contract', async () => {
        await truffleAssert.reverts(
            bsh_coreV2.updateBSHPeriphery(accounts[2], {from: accounts[1]}),
            "Unauthorized"
        );
    });

    it('Scenario 6: Should allow contract owner to update a new URI', async () => {
        var new_uri = 'https://1234.iconee/'
        await bsh_coreV2.updateUri(new_uri);
    });

    it('Scenario 7: Should revert when arbitrary client update a new URI', async () => {
        var new_uri = 'https://1234.iconee/'
        await truffleAssert.reverts(
            bsh_coreV2.updateUri(new_uri, {from: accounts[1]}),
            "Unauthorized"
        );
    });

    it('Scenario 8: Should allow contract owner to update fee ratio', async () => {
        var new_fee = 20;
        await bsh_coreV2.setFeeRatio(new_fee);
    });

    it('Scenario 9: Should revert when an arbitrary client updates fee ratio', async () => {
        var new_fee = 20;
        await truffleAssert.reverts(
            bsh_coreV2.setFeeRatio(new_fee, {from: accounts[1]}),
            "Unauthorized"
        );
    });

    it('Scenario 10: Should revert when Fee Numerator is higher than Fee Denominator', async () => {
        var new_fee = 20000;
        await truffleAssert.reverts(
            bsh_coreV2.setFeeRatio(new_fee),
            "InvalidSetting"
        );
    });

    it('Scenario 11: Should receive an id of a given coin name when querying a valid supporting coin', async () => {
        var _name1 = "wBTC";    var _name2 = "Ethereum";
        await bsh_coreV2.register(_name1);
        await bsh_coreV2.register(_name2);

        var _query = "ICON";
        var id = web3.utils.keccak256(_query);
        var result = await bsh_coreV2.coinId(_query);
        assert(
            web3.utils.BN(result).toString() === web3.utils.toBN(id).toString()
        );
    }); 

    it('Scenario 12: Should receive an id = 0 when querying an invalid supporting coin', async () => {
        var _query = "EOS";
        var result = await bsh_coreV2.coinId(_query);
        assert(
            web3.utils.BN(result).toNumber() === 0
        );
    }); 

    it('Scenario 13: Should revert when a non-Owner tries to add a new Owner', async () => {
        var oldList = await bsh_coreV2.getOwners();
        await truffleAssert.reverts(
            bsh_coreV2.addOwner(accounts[1], {from: accounts[2]}),
            "Unauthorized"
        );
        var newList = await bsh_coreV2.getOwners();
        assert(
            oldList.length === 1 && oldList[0] === accounts[0] &&
            newList.length === 1 && newList[0] === accounts[0]
        );
    }); 

    it('Scenario 14: Should allow a current Owner to add a new Owner', async () => {
        var oldList = await bsh_coreV2.getOwners();
        await bsh_coreV2.addOwner(accounts[1]);
        var newList = await bsh_coreV2.getOwners();
        assert(
            oldList.length === 1 && oldList[0] === accounts[0] &&
            newList.length === 2 && newList[0] === accounts[0] && newList[1] === accounts[1]
        );
    }); 
    
    it('Scenario 15: Should allow old owner to register a new coin - After adding new Owner', async () => {
        var _name3 = "TRON";
        await bsh_coreV2.register(_name3);
        output = await bsh_coreV2.coinNames();
        assert(
            output[0] === _native && output[1] === 'ICON' &&
            output[2] === 'wBTC' && output[3] === 'Ethereum' &&
            output[4] ===  'TRON'
        );
    });

    it('Scenario 16: Should allow new owner to register a new coin', async () => {   
        var _name3 = "BINANCE";
        await bsh_coreV2.register(_name3, {from: accounts[1]});
        output = await bsh_coreV2.coinNames();
        assert(
            output[0] === _native && output[1] === 'ICON' &&
            output[2] === 'wBTC' && output[3] === 'Ethereum' &&
            output[4] ===  'TRON' && output[5] === 'BINANCE'
        );
    }); 

    it('Scenario 17: Should allow new owner to update BSHPeriphery contract', async () => {
        await bsh_coreV2.updateBSHPeriphery(accounts[3], {from: accounts[1]});
    });

    it('Scenario 18: Should also allow old owner to update BSHPeriphery contract - After adding new Owner', async () => {
        await bsh_coreV2.updateBSHPeriphery(accounts[3], {from: accounts[0]});
    });

    it('Scenario 19: Should allow new owner to update the new URI', async () => {
        var new_uri = 'https://1234.iconee/'
        await bsh_coreV2.updateUri(new_uri, {from: accounts[1]});
    });

    it('Scenario 20: Should also allow old owner to update the new URI - After adding new Owner', async () => {
        var new_uri = 'https://1234.iconee/'
        await bsh_coreV2.updateUri(new_uri, {from: accounts[0]});
    });

    it('Scenario 21: Should allow new owner to update new fee ratio', async () => {
        var new_fee = 30;
        await bsh_coreV2.setFeeRatio(new_fee, {from: accounts[1]});
    });

    it('Scenario 22: Should also allow old owner to update new fee ratio - After adding new Owner', async () => {
        var new_fee = 30;
        await bsh_coreV2.setFeeRatio(new_fee, {from: accounts[0]});
    });

    it('Scenario 23: Should revert when non-Owner tries to remove an Owner', async () => {
        var oldList = await bsh_coreV2.getOwners();
        await truffleAssert.reverts(
            bsh_coreV2.removeOwner(accounts[0], {from: accounts[2]}),
            "Unauthorized"
        );
        var newList = await bsh_coreV2.getOwners();
        assert(
            oldList.length === 2 && oldList[0] === accounts[0] && oldList[1] === accounts[1] &&
            newList.length === 2 && newList[0] === accounts[0] && newList[1] === accounts[1]
        );
    });

    it('Scenario 24: Should allow one current Owner to remove another Owner', async () => {
        var oldList = await bsh_coreV2.getOwners();
        await bsh_coreV2.removeOwner(accounts[0], {from: accounts[1]});
        var newList = await bsh_coreV2.getOwners();
        assert(
            oldList.length === 2 && oldList[0] === accounts[0] && oldList[1] === accounts[1] &&
            newList.length === 1 && newList[0] === accounts[1]
        );
    });

    it('Scenario 25: Should revert when the last Owner removes him/herself', async () => {
        var oldList = await bsh_coreV2.getOwners();
        await truffleAssert.reverts(
            bsh_coreV2.removeOwner(accounts[1], {from: accounts[1]}),
            "Unable to remove last Owner"
        );
        var newList = await bsh_coreV2.getOwners();
        assert(
            oldList.length === 1 && oldList[0] === accounts[1] &&
            newList.length === 1 && newList[0] === accounts[1]
        );
    });

    it('Scenario 26: Should revert when removed Owner tries to register a new coin', async () => {
        var _name3 = "KYBER";
        await truffleAssert.reverts(
            bsh_coreV2.register(_name3),
            'Unauthorized'
        );
        output = await bsh_coreV2.coinNames();
        assert(
            output[0] === _native && output[1] === 'ICON' &&
            output[2] === 'wBTC' && output[3] === 'Ethereum' &&
            output[4] ===  'TRON' && output[5] === 'BINANCE'
        );
    });

    it('Scenario 27: Should revert when removed Owner tries to update BSHPeriphery contract', async () => {
        await truffleAssert.reverts(
            bsh_coreV2.updateBSHPeriphery(accounts[3], {from: accounts[0]}),
            'Unauthorized'
        );
    });

    it('Scenario 28: Should revert when removed Owner tries to update the new URI', async () => {
        var new_uri = 'https://1234.iconee/'
        await truffleAssert.reverts(
            bsh_coreV2.updateUri(new_uri, {from: accounts[0]}),
            'Unauthorized'
        );
    });

    it('Scenario 29: Should revert when removed Owner tries to update new fee ratio', async () => {
        var new_fee = 30;
        await truffleAssert.reverts(
            bsh_coreV2.setFeeRatio(new_fee, {from: accounts[0]}),
            'Unauthorized'
        );
    });

    it('Scenario 30: Should allow arbitrary client to query balance of an account', async () => {
        var _coin = 'ICON';
        var _id = await bsh_coreV2.coinId(_coin);
        var _value = 2000;
        await bsh_coreV2.mintMock(accounts[2], _id, _value);
        var balance = await bsh_coreV2.getBalanceOf(accounts[2], _coin, {from: accounts[2]});
        assert(
            web3.utils.BN(balance._usableBalance).toNumber() === _value &&
            web3.utils.BN(balance._lockedBalance).toNumber() === 0 &&
            web3.utils.BN(balance._refundableBalance).toNumber() === 0
        );
    });

    it('Scenario 31: Should allow arbitrary client to query a batch of balances of an account', async () => {
        var _coin1 = 'ICON';        var _coin2 = 'TRON';             
        var _id = await bsh_coreV2.coinId(_coin2);
        var _value = 10000;         var another_value = 2000;
        await bsh_coreV2.mintMock(accounts[2], _id, _value);
        var balance = await bsh_coreV2.getBalanceOfBatch(accounts[2], [_coin1,_coin2], {from: accounts[2]});
        assert(
            web3.utils.BN(balance._usableBalances[0]).toNumber() === another_value &&
            web3.utils.BN(balance._lockedBalances[0]).toNumber() === 0 &&
            web3.utils.BN(balance._refundableBalances[0]).toNumber() === 0 &&

            web3.utils.BN(balance._usableBalances[1]).toNumber() === _value &&
            web3.utils.BN(balance._lockedBalances[1]).toNumber() === 0 &&
            web3.utils.BN(balance._refundableBalances[1]).toNumber() === 0
        );
    });

    it('Scenario 32: Should allow arbitrary client to query an Accumulated Fees', async () => {
        var _coin1 = 'ICON';    var _coin2 = 'TRON';
        var _value1 = 40000;    var _value2 = 10000;    var _native_value = 2000;
        await bsh_coreV2.setAggregationFee(_coin1, _value1);
        await bsh_coreV2.setAggregationFee(_coin2, _value2);
        await bsh_coreV2.setAggregationFee(_native, _native_value);
        var fees = await bsh_coreV2.getAccumulatedFees({from: accounts[3]});
        assert(
            fees[0].coinName === _native && Number(fees[0].value) === _native_value &&
            fees[1].coinName === _coin1 && Number(fees[1].value) === _value1 &&
            fees[2].coinName === 'wBTC' && Number(fees[2].value) === 0 &&
            fees[3].coinName === 'Ethereum' && Number(fees[3].value) === 0 &&
            fees[4].coinName === _coin2 && Number(fees[4].value) === _value2
        );
    });

    it('Scenario 33: Should revert when a client reclaims an exceeding amount', async () => {
        var _coin = 'ICON';     var _value = 10000;     var _exceedAmt = 20000;
        await bsh_coreV2.setRefundableBalance(accounts[2], _coin, _value);
        await truffleAssert.reverts(
            bsh_coreV2.reclaim(_coin, _exceedAmt, {from: accounts[2]}),
            "Imbalance"
        );
    });

    it('Scenario 34: Should revert when a client, which does not own a refundable, tries to reclaim', async () => {
        var _coin = 'ICON';     var _value = 10000;
        await truffleAssert.reverts(
            bsh_coreV2.reclaim(_coin, _value, {from: accounts[3]}),
            "Imbalance"
        );
    });

    it('Scenario 35: Should succeed when a client, which owns a refundable, tries to reclaim', async () => {
        var _coin = 'ICON';     var _value = 10000;
        var _id = await bsh_coreV2.coinId(_coin);
        await bsh_coreV2.mintMock(bsh_coreV2.address, _id, _value);
        var balanceBefore = await bsh_coreV2.getBalanceOf(accounts[2], _coin);
        await bsh_coreV2.reclaim(_coin, _value, {from: accounts[2]});
        var balanceAfter = await bsh_coreV2.getBalanceOf(accounts[2], _coin);
        assert(
            web3.utils.BN(balanceAfter._usableBalance).toNumber() === 
                web3.utils.BN(balanceBefore._usableBalance).toNumber() + _value
        );
    });

    it('Scenario 36: Should not allow any clients (even a contract owner) to call a refund()', async () => {
        //  This function should be called only by itself (BSHCore contract)
        var _coin = 'ICON';     var _value = 10000;
        await truffleAssert.reverts(
            bsh_coreV2.refund(accounts[0], _coin, _value),
            "Unauthorized"
        );
    });

    it('Scenario 37: Should not allow any clients (even a contract owner) to call a mint()', async () => {
        //  This function should be called only by BSHPeriphery contract
        var _coin = 'ICON';     var _value = 10000;
        await truffleAssert.reverts(
            bsh_coreV2.mint(accounts[0], _coin, _value),
            "Unauthorized"
        );
    });

    it('Scenario 38: Should not allow any clients (even a contract owner) to call a handleResponseService()', async () => {
        //  This function should be called only by BSHPeriphery contract
        var RC_OK = 0;
        var _coin = 'ICON';     var _value = 10000;     var _fee = 1;   var _rspCode = RC_OK;
        await truffleAssert.reverts(
            bsh_coreV2.handleResponseService(accounts[0], _coin, _value, _fee, _rspCode),
            "Unauthorized"
        );
    });

    it('Scenario 39: Should not allow any clients (even a contract owner) to call a handleErrorFeeGathering()', async () => {
        //  This function should be called only by BSHPeriphery contract
        var _fees = [ [_native, 1000], ['ICON', 1000], ['TRON', 1000], ['Ethereum', 1000]];
        await truffleAssert.reverts(
            bsh_coreV2.handleErrorFeeGathering(_fees),
            "Unauthorized"
        );
    });

    it('Scenario 40: Should not allow any clients (even a contract owner) to call a gatherFeeRequest()', async () => {
        //  This function should be called only by BSHPeriphery contract
        await truffleAssert.reverts(
            bsh_coreV2.gatherFeeRequest(),
            "Unauthorized"
        );
    });

    it('Scenario 41: Should update AggregationFee state variable when BSHPeriphery calls handleErrorFeeGathering()', async () => {
        //  This function should be called only by BSHPeriphery contract
        //  In this test, I'm using accounts[2] as a BSHPeriphery contract
        await bsh_coreV2.clearAggregationFee();   // this step clear all aggregation fees were set before
        var _fees = [ [_native, 1000], ['ICON', 1000], ['TRON', 1000], ['Ethereum', 1000]];
        var oldFees = await bsh_coreV2.getAccumulatedFees();
        await bsh_coreV2.updateBSHPeriphery(accounts[2], {from: accounts[1]});
        await bsh_coreV2.handleErrorFeeGathering(_fees, {from: accounts[2]});
        var updatedFees = await bsh_coreV2.getAccumulatedFees();
        assert(
            oldFees[0].coinName === _native && Number(oldFees[0].value) === 0 &&
            oldFees[1].coinName === 'ICON' && Number(oldFees[1].value) === 0 &&
            oldFees[2].coinName === 'wBTC' && Number(oldFees[2].value) === 0 &&
            oldFees[3].coinName === 'Ethereum' && Number(oldFees[3].value) === 0 &&
            oldFees[4].coinName === 'TRON' && Number(oldFees[4].value) === 0 &&

            updatedFees[0].coinName === _native && Number(updatedFees[0].value) === 1000 &&
            updatedFees[1].coinName === 'ICON' && Number(updatedFees[1].value) === 1000 &&
            updatedFees[2].coinName === 'wBTC' && Number(updatedFees[2].value) === 0 &&
            updatedFees[3].coinName === 'Ethereum' && Number(updatedFees[3].value) === 1000 &&
            updatedFees[4].coinName === 'TRON' && Number(updatedFees[4].value) === 1000
        );
    });

    it('Scenario 42: Should return an array of charging fees when BSHPeriphery calls gatherFeeRequest()', async () => {
        //  This function should be called only by BSHPeriphery contract
        //  In this test, I'm using accounts[2] as a BSHPeriphery contract
        var result = await bsh_coreV2.gatherFeeRequest.call({from: accounts[2]});
        assert(
            result[0].coinName === _native && Number(result[0].value) === 1000 &&
            result[1].coinName === 'ICON' && Number(result[1].value) === 1000 &&
            result[2].coinName === 'Ethereum' && Number(result[2].value) === 1000 &&
            result[3].coinName === 'TRON' && Number(result[3].value) === 1000
        );
    });
});