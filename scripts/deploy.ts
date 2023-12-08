import { toNano } from 'locklift';

// npx locklift run --config locklift.config.ts --network test --script scripts/deploy.ts

async function main() {
    const signer = await locklift.keystore.getSigner('0');

    const { contract: testContract } = await locklift.factory.deployContract({
        // name of your contract
        contract: 'Test',
        // runtime deployment arguments
        constructorParams: {
            data: "data"
        },
        // static parameters of contract
        initParams: {},
        // public key in init data
        publicKey: signer!.publicKey,
        // this value will be transfered from giver to deployable contract
        value: toNano(1),
    });

    console.log(testContract.address.toString());
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e);
        process.exit(1);
    });

