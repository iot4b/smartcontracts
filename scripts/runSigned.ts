import {Address} from 'locklift';
import {data} from './data';

// npx locklift run --config locklift.config.ts --network main --script scripts/runSigned.ts

async function main() {
    await run(data.Device.KeeneticExtra)
    await run(data.Device.KeeneticOmni)
}

async function run(inst: any) {
    console.log(`Get ${inst.name} contract at address ${inst.address}`)
    locklift.keystore.addKeyPair(inst.keyPair)

    const contract = locklift.factory.getDeployedContract(inst.name, new Address(inst.address))

    console.log('setDeviceAPI...')
    const res = await contract.methods
        .setDeviceAPI({deviceAPI: data.DeviceAPI.Keenetic.address})
        .sendExternal({publicKey: inst.keyPair.publicKey})
    console.log(res.output)
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
