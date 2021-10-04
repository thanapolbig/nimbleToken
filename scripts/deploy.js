async function main() {
    [owner, addr1, addr2, ...addrs] = await ethers.getSigners();
    //deploy any evn
    const NimbleToken = await ethers.getContractFactory("NimbleToken",owner);
    const SyrupBar = await ethers.getContractFactory("SyrupBar",owner);
    const EventBar = await ethers.getContractFactory("EventBar",owner);
    const MasterChef = await ethers.getContractFactory("MasterChef",owner);

    let nimbleToken = await NimbleToken.deploy();
    let syrupBar = await SyrupBar.deploy(nimbleToken.address);
    let eventBar = await EventBar.deploy(nimbleToken.address);
    let masterChef = await MasterChef.deploy(nimbleToken.address,syrupBar.address,eventBar.address);

    await nimbleToken.transferOwnership(masterChef.address);
    await syrupBar.transferOwnership(masterChef.address);
    await eventBar.transferOwnership(masterChef.address);

    console.log("nimbleToken deployed to:", nimbleToken.address);
    console.log("syrupBar deployed to:", syrupBar.address);
    console.log("eventBar deployed to:", eventBar.address);
    console.log("masterChef deployed to:", masterChef.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });