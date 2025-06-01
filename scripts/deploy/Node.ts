import {toNano} from 'locklift';
import {generateSignKeys} from '../util';
import {data} from '../data';

// npx locklift run --config locklift.config.ts --network main --script scripts/deploy/Node.ts

async function main() {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying Node Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'Node',
        constructorParams: {
            elector: data.Elector.address,
            ipPort: '127.0.0.1',
            contactInfo: '',
        },
        publicKey: signer.publicKey,
        value: toNano(1),
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

