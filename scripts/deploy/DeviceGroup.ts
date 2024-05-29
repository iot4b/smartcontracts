import { toNano } from 'locklift';
import { generateSignKeys } from '../util';

// npx locklift run --config locklift.config.ts --network test --script scripts/deploy/DeviceGroup.ts

async function main() {
    const signer = await generateSignKeys()
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying DeviceGroup Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'DeviceGroup',
        constructorParams: {
            name: 'Test Device Group',
            elector: '0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a',
            owners: [
                [
                    '0x6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b',
                    '0:0000000000000000000000000000000000000000000000000000000000000000'
                ]
            ],
            devices: [
                [
                    '0:306824cc3dbec3e8f1f249a61ef111e0d66034063dc385aa77b52b4d02bfd68d',
                    '0xa6493ea99cfc7ed521eabc8f6b4afbd83d313d14e30002cc184d33335f1e09d7'
                ],
                [
                    '0:c292c8d2fe796d62671ff5fd21ae544481bf020039c0626736604b587db1058a',
                    '0x8527e6475d80ee688fbf84bc1c54074fe51f037d9fc683c32e31b74341131878'
                ],
                [
                    '0:09a02fd8865682527ce000c94e268eba6a0a192fea462ecdc7341855ecd931e6',
                    '0x5389801f7c4c9a47819085d8a18aebe5f7a51b810aa3e321d05a9790670e158d'
                ]
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

