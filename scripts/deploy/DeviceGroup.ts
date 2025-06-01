import {toNano} from 'locklift';
import {generateSignKeys} from '../util';
import {data} from '../data';

// npx locklift run --config locklift.config.ts --network main --script scripts/deploy/DeviceGroup.ts

async function main() {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying DeviceGroup Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'DeviceGroup',
        constructorParams: {
            name: 'Test Device Group',
            elector: data.Elector.address,
            owners: [
                [
                    `0x${data.DeviceGroup.keyPair.publicKey}`,
                    '0:0000000000000000000000000000000000000000000000000000000000000000'
                ]
            ],
            devices: [
                [
                    data.Device.Device1.address,
                    `0x${data.Device.Device1.keyPair.publicKey}`,
                ],
                [
                    data.Device.Device2.address,
                    `0x${data.Device.Device2.keyPair.publicKey}`,
                ],
                [
                    data.Device.Device3.address,
                    `0x${data.Device.Device3.keyPair.publicKey}`,
                ]
            ]
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

