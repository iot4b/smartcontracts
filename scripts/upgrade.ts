import {Address} from 'locklift';

// npx locklift run --config locklift.config.ts --network test --script scripts/upgrade.ts

async function main() {
    const device = 'Device'
    await upgrade(device, '0:260daa32cd40ba0bd4f803cff1896f51fead885c1a292ab46f6a02d771c84b37')
    await upgrade(device, '0:ca2aea5c83b0a04c70eb6bd57fbcd4974428a9f2758e8748b050d0b563fc8819')
    await upgrade(device, '0:03b678d0869a6d522e9e42d79d4527655ae920693528df186002b0d026f0b453')

    const node = 'Node'
    await upgrade(node, '0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26')
    await upgrade(node, '0:86429800dd5b8ddc9a1283341b106cdb7acb2807c4e5f91e523c2803e6c76ddd')
    await upgrade(node, '0:e986b8305e5d46cc221cc9e14785bfe361b8558104396bdc082fa4c6321ffc68')
}

async function upgrade(name: string, address: string) {
    console.log(`Upgrading ${name} contract at address ${address}`)

    const contract = locklift.factory.getDeployedContract(name, new Address(address))
    const artifacts = locklift.factory.getContractArtifacts(name)

    const res = await contract.methods.upgrade({
        code: artifacts.code,
    }).sendExternal({withoutSignature: true})

    console.log(res.transaction.exitCode ? `Exit code: ${res.transaction.exitCode}` : 'Success!')
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
