import {toNano} from 'locklift';
import {generateSignKeys} from '../util';
import {data} from '../data';

// npx locklift run --config locklift.config.ts --network main --script scripts/deploy/Device.ts

async function main() {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying Device Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'Device',
        constructorParams: {
            name: 'Test Device',
            elector: data.Elector.address,
            vendor: data.Vendor.IOT4Linux.address,
            group: data.DeviceGroup.address,
            owners: [
                [
                    `0x${data.DeviceGroup.keyPair.publicKey}`,
                    '0:0000000000000000000000000000000000000000000000000000000000000000'
                ]
            ],
            dtype: 'device emulator',
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

