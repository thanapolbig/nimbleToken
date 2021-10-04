const { expect } = require("chai");
const { ethers } = require("hardhat");
// import NimbleToken from '../artifacts/contracts/NimbleToken.sol/NimbleToken.json'

describe("NimbleToken", function(){

    let NimbelToken;
    let nimbleToken;
    let owner;
    let addr1;
    let addr2;
    let addrs;

    beforeEach(async function () {
        // Get the ContractFactory and Signers here.
        NimbelToken = await ethers.getContractFactory("NimbleToken",owner);
        [owner, addr1, addr2, ...addrs] = await ethers.getSigners();

        nimbleToken = await NimbelToken.deploy();
    });

    it("decimals,name,symbol", async function(){
        // [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
        // const NimbleToken = await ethers.getContractFactory("NimbleToken",owner);
        // const nimbleToken = await NimbleToken.deploy();
        // await nimbleToken.deployed();
        expect(await nimbleToken.decimals()).to.equal(3);
        expect(await nimbleToken.name()).to.equal("Nimble Token");
        expect(await nimbleToken.symbol()).to.equal("Nimble");
        expect(await nimbleToken.owner()).to.equal(owner.address);
    });

    it("mint", async function(){
        await nimbleToken.mint(owner.address,50000);
        expect(await nimbleToken.totalSupply()).to.equal(50000);
        expect(await nimbleToken.balanceOf(owner.address)).to.equal(50000);

        // expect(await nimbleToken.connect(addr1).mint(addr1.address,50000)).to.be.revertedWith("Ownable: caller is not the owner");
        // expect(await nimbleToken.balanceOf(addr1.address)).to.equal(50000);

    });

    it("approve,allowance", async function () {
        await nimbleToken.connect(owner).approve(addr1.address,50000);
        expect(await nimbleToken.allowance(owner.address,addr1.address)).to.equal(50000);
    });

    it("transfer", async function () {
        await nimbleToken.connect(owner).mint(owner.address,50000);
        await nimbleToken.connect(owner).transfer(addr1.address,10000);
        const ownerBalance = await nimbleToken.balanceOf(owner.address)
        const addr1Balance = await nimbleToken.balanceOf(addr1.address)
        expect(ownerBalance).to.equal(40000);
        expect(addr1Balance).to.equal(10000);
    });

    it("transferFail", async function () {
        // await nimbleToken.connect(owner).transfer(addr1.address,10000);
        await expect(
            nimbleToken.connect(addr1).transfer(owner.address, 1)
        ).to.be.revertedWith("ERC20: transfer amount exceeds balance");
    });

    it("transferFrom", async function () {
        await nimbleToken.connect(owner).mint(owner.address,50000);
        await nimbleToken.connect(owner).approve(addr1.address,30000);
        await nimbleToken.connect(addr1).transferFrom(owner.address,addr1.address,30000)
        expect(await nimbleToken.balanceOf(owner.address)).to.equal(20000);
        expect(await nimbleToken.balanceOf(addr1.address)).to.equal(30000);
    });

    it("transferFromFail", async function () {
        await nimbleToken.connect(owner).mint(owner.address,50000);
        await nimbleToken.connect(owner).approve(addr1.address,30000);
        await expect (
            nimbleToken.connect(addr1).transferFrom(owner.address,addr1.address,40000)
        ).to.be.revertedWith("ERC20: transfer amount exceeds allowance");
        const ownerBalance = await nimbleToken.balanceOf(owner.address)
        const addr1Balance = await nimbleToken.balanceOf(addr1.address)
        expect(ownerBalance).to.equal(50000);
        expect(addr1Balance).to.equal(0);
    });
});