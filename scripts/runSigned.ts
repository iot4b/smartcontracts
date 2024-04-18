import {Address} from 'locklift';
import {data} from "./data";

// npx locklift run --config locklift.config.ts --network test --script scripts/runSigned.ts

async function main() {
    locklift.keystore.addKeyPair(data.testDevice1.keyPair)
    await run(data.testDevice1)
}

async function run(data: any) {
    console.log(`Get ${data.name} contract at address ${data.address}`)

    const contract = locklift.factory.getDeployedContract(data.name, new Address(data.address))
    const res = await contract.methods
        .setLock({lock: true})
        .sendExternal({publicKey: data.keyPair.publicKey})

    console.log(res.output)
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
