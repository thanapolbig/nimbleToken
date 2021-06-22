// We require the Hardhat Runtime Environment explicitly here. This is optional 
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const wallet = "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"
const wallet2 = "0x70997970c51812dc3a010c7d01b50e0d17dc79c8"



async function main() {
  // Hardhat always runs the compile task when running scripts with its command
  // line interface.
  //
  // If this script is run directly using `node` you may want to call compile 
  // manually to make sure everything is compiled
  // await hre.run('compile');

  // We get the contract to deploy

  //--Get wallet account form hardhat
  const [owner, addr1] = await ethers.getSigners();
  // console.log(addr1.address);

  //--Call smart contract
  const SimpleBank = await hre.ethers.getContractFactory("SimpleBank");
  const bank = await SimpleBank.deploy();
  await bank.deployed();

  const enroll = await bank.connect(owner).enroll("big");

  const enroll2 = await bank.connect(addr1).enroll("ohm");

  const deposit = await bank.connect(owner).deposit(5000);

  const withdraw = await bank.connect(owner).withdraw(1000);

  const TransferTo = await bank.connect(owner).TransferTo(addr1.address,1000)

  const wallet = await bank.wallet("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266")
  const wallet2 = await bank.wallet("0x70997970c51812dc3a010c7d01b50e0d17dc79c8")
  
  var outputWallet1 = {
    Name : wallet[0],
    Balances : parseInt(wallet[1]),
    owner : wallet[2]
  }
  var outputWallet2 = {
    Name : wallet2[0],
    Balances : parseInt(wallet2[1]),
    owner : wallet2[2]
  }
  console.log("walletA : " ,outputWallet1);
  console.log("walletB : " ,outputWallet2);


  //Hexadecimal => Decimal
  // console.log("Wallet1 : ",  await bank.wallet("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"));
  // console.log("Wallet2 : ",  await bank.wallet("0x70997970c51812dc3a010c7d01b50e0d17dc79c8"));
  
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
