import {toNano} from 'locklift';
import {generateSignKeys} from '../util';
import {data} from '../data';

// npx locklift run --config locklift.config.ts --network main --script scripts/deploy/DeviceAPI.ts

async function main() {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);

    console.log('Deploying DeviceAPI Contract...');
    const { contract } = await locklift.factory.deployContract({
        contract: 'DeviceAPI',
        constructorParams: {
            elector: data.Elector.address,
            vendor: data.Vendor.IOT4Linux.address,
            deviceType: 'linux',
            deviceModel: '',
            deviceFW: '0.0.1',
            deviceName: 'iot4b-device',
            deviceConfig: '{"listCard": {"title": "{{name}}", "subtitle": "{{vendor}}, {{model}}"}, "cardParams": [{"title": "Vendor", "value": "{{vendorName}}", "type": "address"}, {"title": "Node", "value": "{{node}}", "type": "address"}, {"title": "Firmware", "value": "{{version}}", "type": "string"}, {"title": "Registration", "value": "{{lastRegisterTime}}", "type": "datetime"}], "cardCharts": [], "cardButtons": [{"title": "Device Info", "dialog": null, "cmd": "uname -a", "displayResult": "text"}, {"title": "Reboot", "dialog": null, "cmd": "reboot", "displayResult": "ok"}]}'
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

