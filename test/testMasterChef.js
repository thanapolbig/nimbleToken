const { expect } = require("chai");
const { ethers } = require("hardhat");
// const { Contract } =  require("ethers") ;

const { accounts, contract } = require('@openzeppelin/test-helpers');

describe("MasterChef", function(){
    let nimbleToken;
    let syrupBar;
    let eventBar;
    let nimbleTokenId;
    let syrupBarId;
    let eventBarId;
    let MasterChef;
    let masterChef;
    let masterChefId;
    let owner;
    let addr1;
    let addr2;
    let addrs;

    beforeEach(async function () {
        [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
        //deploy any evn
        const NimbleToken = await ethers.getContractFactory("NimbleToken",owner);
        const SyrupBar = await ethers.getContractFactory("SyrupBar",owner);
        const EventBar = await ethers.getContractFactory("EventBar",owner);
        //deploy nimbleToken
        nimbleToken = await NimbleToken.deploy();
        nimbleTokenId = nimbleToken.address
        //deploy syrupBar,EventBar
        syrupBar = await SyrupBar.deploy(nimbleTokenId);
        eventBar = await EventBar.deploy(nimbleTokenId);
        syrupBarId = syrupBar.address
        eventBarId = eventBar.address
        // console.log(nimbleTokenId);
        // console.log(syrupBarId);
        // console.log(eventBarId);
        //deploy masterChef
        MasterChef = await ethers.getContractFactory("MasterChef",owner);
        masterChef = await MasterChef.deploy(nimbleTokenId,syrupBarId,eventBarId);
        masterChefId = masterChef.address;
    });

    it("owner", async function(){
        expect(await masterChef.owner()).to.equal(owner.address);
        expect(await masterChef.nimble()).to.equal(nimbleTokenId);
        expect(await masterChef.syrup()).to.equal(syrupBarId);
        expect(await masterChef.eventbar()).to.equal(eventBarId);
    });

    it("transferOwnership", async function(){
        await nimbleToken.transferOwnership(masterChefId);
        await syrupBar.transferOwnership(masterChefId);
        await eventBar.transferOwnership(masterChefId);
        expect(await nimbleToken.owner()).to.equal(masterChefId);
        expect(await syrupBar.owner()).to.equal(masterChefId);
        expect(await eventBar.owner()).to.equal(masterChefId);
    });

    it("mint", async function(){
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint();
        const totalSupply = await nimbleToken.totalSupply();
        const bal = await nimbleToken.balanceOf(syrupBarId);
        expect(ethers.utils.formatUnits(totalSupply,3)).to.equal("50000.0");
        expect(ethers.utils.formatUnits(bal,3)).to.equal("50000.0");
    });

    it("autoClaimCheckin", async () => {
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint(); //50000
        // 20 % => 10000
        await masterChef.connect(owner).addWorkday(200); // 10000/200 = 50token/day
        const addressArray = [owner.address,addr1.address,addr2.address,addrs[0].address,addrs[1].address]
        await masterChef.connect(owner).autoClaimCheckin(addressArray) // 10 /address
        const balOwner = await nimbleToken.balanceOf(owner.address);
        const balAddr1 = await nimbleToken.balanceOf(addr1.address);
        const balAddr2 = await nimbleToken.balanceOf(addr2.address);
        const balAddr3 = await nimbleToken.balanceOf(addrs[0].address);
        const balAddr4 = await nimbleToken.balanceOf(addrs[1].address);
        expect(ethers.utils.formatUnits(balOwner,3)).to.equal("10.0")
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("10.0")
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("10.0")
        expect(ethers.utils.formatUnits(balAddr3,3)).to.equal("10.0")
        expect(ethers.utils.formatUnits(balAddr4,3)).to.equal("10.0")
    })

    it("vote",async ()=>{
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint(); //50000
        const addressArray = [owner.address,addr1.address]
        await masterChef.connect(owner).addVote(addressArray);
        expect(await masterChef.getRightScore(owner.address)).to.equal("4");
        await masterChef.connect(owner).vote(addr1.address,"2");    //addr1 : 2
        await masterChef.connect(addr1).vote(owner.address,"3");    //owner : 3
        expect(await masterChef.getScore(addr1.address)).to.equal("2");
        expect(await masterChef.getRightScore(addr1.address)).to.equal("1");
        expect(await masterChef.getScore(owner.address)).to.equal("3");
        expect(await masterChef.getRightScore(owner.address)).to.equal("2");
        await masterChef.connect(owner).claimRewardScoreVote(addressArray); // 5/10000 = 1000/1score
        const balOwner = await nimbleToken.balanceOf(owner.address);
        const balAddr1 = await nimbleToken.balanceOf(addr1.address);
        // console.log(ethers.utils.formatEther(balAddr1))
        expect(ethers.utils.formatUnits(balOwner,3)).to.equal("6000.0")  //6000
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("4000.0")  //4000
    })

    it("eventByUser", async () => {
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint(); //50000
        await masterChef.connect(owner).addWorkday(10); // 10000/10 = 1000/day
        const addressArray = [addr1.address,addr2.address]
        await masterChef.connect(owner).autoClaimCheckin(addressArray) // 500/address
        let balAddr1 = await nimbleToken.balanceOf(addr1.address);
        let balAddr2 = await nimbleToken.balanceOf(addr2.address);
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("500.0")
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("500.0")

        // --------allowance-----------
        await nimbleToken.connect(addr1).approve(eventBarId,ethers.utils.parseUnits('500',3)) //500
        const alloAddr1 = await nimbleToken.connect(addr1).allowance(addr1.address,eventBarId)
        expect(ethers.utils.formatUnits(alloAddr1,3)).to.equal("500.0")

        //--------- create Event ---------
        const reward = ethers.utils.parseUnits('250',3)
        await masterChef.connect(addr1).createEvent("dota2","play dota2",reward,Date.now()-3600000)
        let arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[0]).to.equal(addr1.address)
        expect(arrayEvent[1]).to.equal("dota2")
        expect(arrayEvent[2]).to.equal("play dota2")
        expect(ethers.utils.formatUnits(arrayEvent[3],3)).to.equal("250.0")
        //check money
        balAddr1 = await nimbleToken.balanceOf(addr1.address);
        let balEventBar = await nimbleToken.balanceOf(eventBarId);
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("250.0")
        expect(ethers.utils.formatUnits(balEventBar,3)).to.equal("250.0")

        //-------- startEvent ---------
        await network.provider.send("evm_setNextBlockTimestamp", [Date.now()]) //set timestamp in block
        await masterChef.connect(addr1).startEvent(0)
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("1")

        //--------- joinEvent ---------
        await masterChef.connect(addr2).joinEvent(0)
        await masterChef.connect(addrs[0]).joinEvent(0)

        //----------- AcceptEvent ---------
        const listAddr = [addr2.address,addrs[0].address]
        await masterChef.connect(addr1).AcceptEvent(0,listAddr)
        balAddr2 = await nimbleToken.balanceOf(addr2.address);
        let balAddr3 = await nimbleToken.balanceOf(addrs[0].address);
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("625.0")
        expect(ethers.utils.formatUnits(balAddr3,3)).to.equal("125.0")
        balEventBar = await nimbleToken.balanceOf(eventBarId);
        expect(ethers.utils.formatUnits(balEventBar,3)).to.equal("0.0")
        // event info
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("3")

    })

    it("eventCloseByUser", async () => {
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint(); //50000
        await masterChef.connect(owner).addWorkday(10); // 10000/10 = 1000/day
        const addressArray = [addr1.address,addr2.address]
        await masterChef.connect(owner).autoClaimCheckin(addressArray) // 500/address
        let balAddr1 = await nimbleToken.balanceOf(addr1.address);
        let balAddr2 = await nimbleToken.balanceOf(addr2.address);
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("500.0")
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("500.0")

        // --------allowance-----------
        await nimbleToken.connect(addr1).approve(eventBarId,ethers.utils.parseUnits('500',3)) //500
        const alloAddr1 = await nimbleToken.connect(addr1).allowance(addr1.address,eventBarId)
        expect(ethers.utils.formatUnits(alloAddr1,3)).to.equal("500.0")

        //--------- create Event ---------
        const reward = ethers.utils.parseUnits('250',3)
        await masterChef.connect(addr1).createEvent("dota2","play dota2",reward,Date.now()-3600000)
        let arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[0]).to.equal(addr1.address)
        expect(arrayEvent[1]).to.equal("dota2")
        expect(arrayEvent[2]).to.equal("play dota2")
        expect(ethers.utils.formatUnits(arrayEvent[3],3)).to.equal("250.0")
        //check money
        balAddr1 = await nimbleToken.balanceOf(addr1.address);
        let balEventBar = await nimbleToken.balanceOf(eventBarId);
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("250.0")
        expect(ethers.utils.formatUnits(balEventBar,3)).to.equal("250.0")

        //-------- startEvent ---------
        await network.provider.send("evm_setNextBlockTimestamp", [Date.now()]) //set timestamp in block
        await masterChef.connect(addr1).startEvent(0)
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("1")

        //---------- closeEvent ---------
        await masterChef.connect(addr1).closeEvent(0)
        arrayEvent = await masterChef.eventInfo(0)

        //--------- joinEvent ---------
        await expect (
            masterChef.connect(addr2).joinEvent(0)
        ).to.be.revertedWith("Event can't join")
        await expect (
            masterChef.connect(addrs[0]).joinEvent(0)
        ).to.be.revertedWith("Event can't join")

        //----------- AcceptEvent ---------
        const listAddr = [addr2.address,addrs[0].address]
        await expect(masterChef.connect(addr1).AcceptEvent(0,listAddr)).to.be.revertedWith("status invalid")
        balAddr2 = await nimbleToken.balanceOf(addr2.address);
        let balAddr3 = await nimbleToken.balanceOf(addrs[0].address);
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("500.0")
        expect(ethers.utils.formatUnits(balAddr3,3)).to.equal("0.0")
        balEventBar = await nimbleToken.balanceOf(eventBarId);
        expect(ethers.utils.formatUnits(balEventBar,3)).to.equal("250.0")
        // event info
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("2")

    })


    it("eventByAdmin", async () => {
        await nimbleToken.connect(owner).transferOwnership(masterChefId);
        await syrupBar.connect(owner).transferOwnership(masterChefId);
        await eventBar.connect(owner).transferOwnership(masterChefId);
        await masterChef.connect(owner).mint(); //50000

        //--------- create Event ---------
        const reward = ethers.utils.parseUnits('250',3)
        await masterChef.connect(owner).createEventAdmin("dota2","play dota2",reward,Date.now()-3600000)
        let arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[0]).to.equal(owner.address)
        expect(arrayEvent[1]).to.equal("dota2")
        expect(arrayEvent[2]).to.equal("play dota2")
        expect(ethers.utils.formatUnits(arrayEvent[3],3)).to.equal("250.0")
        //check money
        let balSyrup = await nimbleToken.balanceOf(syrupBarId);
        expect(ethers.utils.formatUnits(balSyrup,3)).to.equal("50000.0")

        //-------- startEvent ---------
        await network.provider.send("evm_setNextBlockTimestamp", [Date.now()]) //set timestamp in block
        await masterChef.connect(owner).startEvent(0)
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("1")

        //--------- joinEvent ---------
        await masterChef.connect(addr1).joinEvent(0)
        await masterChef.connect(addr2).joinEvent(0)

        //----------- AcceptEvent ---------
        let addressArray = [addr1.address,addr2.address]
        await masterChef.connect(owner).AcceptEventAdmin(0,addressArray)
        let balAddr1 = await nimbleToken.balanceOf(addr1.address);
        let balAddr2 = await nimbleToken.balanceOf(addr2.address);
        expect(ethers.utils.formatUnits(balAddr1,3)).to.equal("125.0")
        expect(ethers.utils.formatUnits(balAddr2,3)).to.equal("125.0")
        let balEventBar = await nimbleToken.balanceOf(syrupBarId);
        expect(ethers.utils.formatUnits(balEventBar,3)).to.equal("49750.0")
        // event info
        arrayEvent = await masterChef.eventInfo(0)
        expect(arrayEvent[4]).to.equal("3")
    })




});