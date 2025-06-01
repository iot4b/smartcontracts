import {Address} from 'locklift';
import {data} from './data';

// npx locklift run --config locklift.config.ts --network main --script scripts/elector.ts

async function main() {
    await execute('Elector', data.Elector.address)
}

async function execute(name: string, address: string) {
    console.log(`Get ${name} contract at address ${address}`)

    const contract = locklift.factory.getDeployedContract(name, new Address(address))
    const res = await contract.methods
        .election({})
        // .currentList({})
        // .takeNextRound({
        //     _address: new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26')
        // })
        // .reportFaultNode({
        //     _address: new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26')
        // })
        // .setNodes({
        //     _nodes: [
        //         new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'),
        //         new Address('0:86429800dd5b8ddc9a1283341b106cdb7acb2807c4e5f91e523c2803e6c76ddd'),
        //         new Address('0:e986b8305e5d46cc221cc9e14785bfe361b8558104396bdc082fa4c6321ffc68')
        //     ]
        // })
        .sendExternal({withoutSignature: true})

    console.log(res.output)
}

main()
    .then(() => process.exit(0))
    .catch((e) => {
        console.log(e)
        process.exit(1)
    })
