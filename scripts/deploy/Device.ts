import { toNano } from 'locklift';
import { generateSignKeys } from '../util';

// npx locklift run --config locklift.config.ts --network test --script scripts/deploy/Device.ts

async function main() {
    const signer = await generateSignKeys()
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying Device Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'Device',
        constructorParams: {
            name: 'Test Device',
            elector: '0:20a9657c902ad65dd6eb08b4eb7ec4d3f45bebda5d763bbd7b0b3826824a216a',
            vendor: '0:389e290e6f83d7837f317792c42c370c025c58d34f7e9bdd2c076950acc9c4c9',
            group: '0:0000000000000000000000000000000000000000000000000000000000000000',
            owners: [
                [
                    '0x6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b',
                    '0:0000000000000000000000000000000000000000000000000000000000000000'
                ]
            ],
            dtype: 'dtype',
            version: '0.0.1',
            vendorName: 'vendorName',
            vendorData: 'vendorData',
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

