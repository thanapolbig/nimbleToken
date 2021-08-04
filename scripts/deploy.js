async function main() {
    // We get the contract to deploy
    const Greeter = await ethers.getContractFactory("ERC20");
    const greeter = await Greeter.deploy("big","big");

    console.log("Greeter deployed to:", greeter.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });