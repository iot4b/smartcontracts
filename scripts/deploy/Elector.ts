import { toNano } from 'locklift';
import {generateSignKeys} from "../util";

// npx locklift run --config locklift.config.ts --network test --script scripts/deploy/Elector.ts

async function main() {
    const signer = await generateSignKeys()
    locklift.keystore.addKeyPair(signer);

    const { contract } = await locklift.factory.deployContract({
        contract: 'Elector',
        constructorParams: {
            defaultNodes: [
                '0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26',
                '0:86429800dd5b8ddc9a1283341b106cdb7acb2807c4e5f91e523c2803e6c76ddd',
                '0:e986b8305e5d46cc221cc9e14785bfe361b8558104396bdc082fa4c6321ffc68'
            ]
        },
        publicKey: signer.publicKey,
        value: toNano(10),
    });

    console.log('Deployed to address:');
    console.log(contract.address.toString());
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e);
        process.exit(1);
    });

