import {Address} from 'locklift';

// npx locklift run --config locklift.config.ts --network test --script scripts/upgrade.ts

async function main() {
    const name = 'Device'
    await execute(name, '0:260daa32cd40ba0bd4f803cff1896f51fead885c1a292ab46f6a02d771c84b37')
    await execute(name, '0:ca2aea5c83b0a04c70eb6bd57fbcd4974428a9f2758e8748b050d0b563fc8819')
    await execute(name, '0:03b678d0869a6d522e9e42d79d4527655ae920693528df186002b0d026f0b453')
}

async function execute(name: string, address: string) {
    const contract = locklift.factory.getDeployedContract(name, new Address(address))
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
