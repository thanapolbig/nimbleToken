const { expect } = require("chai");
const { ethers } = require("hardhat");
// import NimbleToken from '../artifacts/contracts/NimbleToken.sol/NimbleToken.json'

describe("NimbleToken", function(){
    before(async function () {
        [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
        this.NimbleToken = await ethers.getContractFactory('NimbleToken',owner);
    });
    beforeEach(async function(){
        this.nimbleToken = await this.NimbleToken.deploy();
        await this.nimbleToken.deployed();
    });

    it("decimals,name,symbol", async function(){
        // [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
        // const NimbleToken = await ethers.getContractFactory("NimbleToken",owner);
        // const nimbleToken = await NimbleToken.deploy();
        // await nimbleToken.deployed();
        expect(await this.nimbleToken.decimals()).to.equal(1);
        expect(await this.nimbleToken.name()).to.equal("Nimble Token");
        expect(await this.nimbleToken.symbol()).to.equal("Nimble");
        expect(await this.nimbleToken.owner()).to.equal(owner.address);
    });

    it("mint", async function(){
        await this.nimbleToken.mint(owner.address,50000);
        expect(await this.nimbleToken.balanceOf(owner.address)).to.equal(50000);

        // expect(await this.nimbleToken.connect(addr1).mint(addr1.address,50000)).to.be.revertedWith("Ownable: caller is not the owner");
        // expect(await this.nimbleToken.balanceOf(addr1.address)).to.equal(50000);

    });

    it("approve,allowance", async function () {
        await this.nimbleToken.connect(owner).approve(addr1.address,50000);
        expect(await this.nimbleToken.allowance(owner.address,addr1.address)).to.equal(50000);
    });

    it("transfer", async function () {
        await this.nimbleToken.connect(owner).mint(owner.address,50000);
        await this.nimbleToken.connect(owner).transfer(addr1.address,10000);
        expect(await this.nimbleToken.balanceOf(owner.address)).to.equal(40000);
        expect(await this.nimbleToken.balanceOf(addr1.address)).to.equal(10000);
    });

    it("transferFrom", async function () {
        await this.nimbleToken.connect(owner).mint(owner.address,50000);
        await this.nimbleToken.connect(owner).approve(addr1.address,30000);
        await this.nimbleToken.connect(addr1).transferFrom(owner.address,addr1.address,30000)
        expect(await this.nimbleToken.balanceOf(owner.address)).to.equal(20000);
        expect(await this.nimbleToken.balanceOf(addr1.address)).to.equal(30000);
    });
});