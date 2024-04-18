import {Address} from 'locklift';
import {data} from './data';

// npx locklift run --config locklift.config.ts --network test --script scripts/execute.ts

async function main() {
    await run(data.testDevice1)
    await run(data.testDevice2)
    await run(data.testDevice3)
}

async function run(data: any) {
    console.log(`Get ${data.name} contract at address ${data.address}`)

    const contract = locklift.factory.getDeployedContract(data.name, new Address(data.address))
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
