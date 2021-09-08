// const { expect } = require("chai");
//
//
// describe("Bank", function() {
//   it("Should return the new greeting once it's changed", async function() {
// //--Get wallet account form hardhat
// const [owner, addr1] = await ethers.getSigners();
// // console.log(addr1.address);
//
// //--Call smart contract
// const SimpleBank = await hre.ethers.getContractFactory("SimpleBank");
// const bank = await SimpleBank.deploy();
// await bank.deployed();
//
// const enroll = await bank.connect(owner).enroll("big");
//
// const enroll2 = await bank.connect(addr1).enroll("ohm");
//
// const deposit = await bank.connect(owner).deposit(5000);
//
// const withdraw = await bank.connect(owner).withdraw(1000);
//
// const TransferTo = await bank.connect(owner).TransferTo(addr1.address,1000)
//
//     expect(await bank.wallet(owner.address)[1]).to.equal(3000);
//
//     expect(await bank.wallet(addr1.address)[1]).to.equal(1000);

//   });
// });
