import {Address} from 'locklift';
import {data} from './data';

// npx locklift run --config locklift.config.ts --network main --script scripts/execute.ts

async function main() {
    await run(data.Device.Device1)
    await run(data.Device.Device2)
    await run(data.Device.Device3)
}

async function run(inst: any) {
    console.log(`Get ${inst.name} contract at address ${inst.address}`)

    const contract = locklift.factory.getDeployedContract(inst.name, new Address(inst.address))
    const res = await contract.methods
        .get({})
        .sendExternal({withoutSignature: true})

    console.log(res.output)
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
