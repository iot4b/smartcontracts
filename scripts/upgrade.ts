import {Address} from 'locklift';
import {data} from './data';

// npx locklift run --config locklift.config.ts --network main --script scripts/upgrade.ts

async function main() {
    await upgrade(data.Device.Device1)
    await upgrade(data.Device.Device2)
    await upgrade(data.Device.Device3)
}

async function upgrade(inst: any) {
    console.log(`Upgrading ${inst.name} contract at address ${inst.address}`)

    const contract = locklift.factory.getDeployedContract(inst.name, new Address(inst.address))
    const artifacts = locklift.factory.getContractArtifacts(inst.name)

    const res = await contract.methods.upgrade({
        code: artifacts.code,
    }).sendExternal({withoutSignature: true})

    console.log(res.transaction.exitCode ? res : 'Success!')
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
