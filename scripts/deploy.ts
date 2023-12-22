import { toNano } from 'locklift';

// npx locklift run --config locklift.config.ts --network test --script scripts/deploy.ts

async function main() {
    const signer = await locklift.keystore.getSigner('0');

    const { contract: testContract } = await locklift.factory.deployContract({
        // name of your contract
        contract: 'Elector',
        // runtime deployment arguments
        constructorParams: {
            defaultNodes: [
                "0:b6f0dd040d1fa3f79871418040e7bb42c02dd674d32cc29fbc82f97dc119a212",
                "0:e413df4ff5e475d6dfd78db61ea635d8eefb18753f3152c51b2929edd684ec57",
                "0:c6bd6af3a4104c9cd3867c79cff5f62410cc95ca94e5e235c8303f9213f19582"
            ]
        },
        // static parameters of contract
        initParams: {},
        // public key in init data
        publicKey: signer!.publicKey,
        // this value will be transfered from giver to deployable contract
        value: toNano(10),
    });

    console.log(testContract.address.toString());
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e);
        process.exit(1);
    });

