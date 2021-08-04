//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.5;

import "hardhat/console.sol";

contract SimpleBank {
    
    struct Wallet{
        string Name;
        uint balances;
        address owner;
    }
    
    mapping (address => Wallet) public wallet;

    
    event LogDepositMade(address indexed accountAddress, uint amount);
    modifier CheckBalances(uint amount){
        require(wallet[msg.sender].balances >= amount, "Insufficient funds");
        _;
    }
    
    function enroll(string memory _Name) public {
        wallet[msg.sender].Name = _Name;
        wallet[msg.sender].balances = 0;
        wallet[msg.sender].owner = msg.sender;
        console.log("init account success");
    }

    
    function deposit(uint depositAmount) public payable {
        wallet[msg.sender].balances += depositAmount;
        emit LogDepositMade(msg.sender, msg.value);
        console.log("deposit success");
    }
    function withdraw(uint withdrawAmount) public CheckBalances(withdrawAmount) {
        wallet[msg.sender].balances -= withdrawAmount;
        console.log("withdraw success");

    }
    
    function TransferTo(address transferTo,uint amount) public CheckBalances(amount) {
            wallet[msg.sender].balances -= amount;
            wallet[transferTo].balances += amount;
        console.log("Transfer success");

    }

}