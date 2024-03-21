import { expect } from "chai";
import { Address, Contract, Signer } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateKeyPair } from 'crypto';

let electorContract: Contract<FactorySource["Elector"]>;
// let signer: Signer;
let publicKey: string;

describe("Elector contract", async function () {
  before(async () => {
    // signer = (await locklift.keystore.getSigner("0"))!;
    // Generate random sign keys
    generateKeyPair('ed25519', {
      publicKeyEncoding: {
        type: 'spki',
        format: 'der'
      },
      privateKeyEncoding: {
        type: 'pkcs8',
        format: 'der',
      }
    }, (err, pub, priv) => { // Callback function
      if (err) {
        console.log("generateKeyPair error: ", err);
      } else {
        publicKey = pub.toString('hex').substring(24);
        const privateKey = priv.toString('hex').substring(32);

        console.log("PublicKey:  ", publicKey);
        console.log("PrivateKey: ", privateKey);

        locklift.keystore.addKeyPair({
          publicKey: publicKey,
          secretKey: privateKey
        });
      }
    });
  });
  describe("Contracts", async function () {
    it("Load contract factory", async function () {
      const sampleData = await locklift.factory.getContractArtifacts("Elector");
    });

    it("Deploy contract", async function () {
      const { contract } = await locklift.factory.deployContract({
        contract: "Elector",
        publicKey: publicKey,
        initParams: {},
        constructorParams: {
          defaultNodes: [],
        },
        value: locklift.utils.toNano(2)
      });
      electorContract = contract;

      // expect(await locklift.provider.getBalance(electorContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Get vendor for vendor", async function () {
      const response = await electorContract.methods.currentList().call();
      expect(response.nodes).to.deep.equal([]);  
    });

    it("Set nodes for elector", async function () {
      const response = await electorContract.methods.setNodes({
        _nodes: [
            new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'),
            new Address('0:86429800dd5b8ddc9a1283341b106cdb7acb2807c4e5f91e523c2803e6c76ddd'),
            new Address('0:e986b8305e5d46cc221cc9e14785bfe361b8558104396bdc082fa4c6321ffc68')
        ]}).call();
        console.log(response);
        const responseCheck = await electorContract.methods.currentList().call();
        expect(responseCheck.nodes).not.to.deep.equal([]);  
    });

    it("Set nodes for elector", async function () {
      const response = await electorContract.methods.takeNextRound({
            _address: new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'),
         }).call();
        console.log(response);
        const responseCheck = await electorContract.methods.participantList().call();
        console.log(responseCheck);
        expect(responseCheck.participants).not.to.deep.equal([]);  
    });
    
    it("Election test of elector", async function () {
      const response = await electorContract.methods.election().call();
        console.log(response);
        const responseCheck = await electorContract.methods.currentList().call();
        console.log(responseCheck);
        expect(responseCheck.nodes).not.to.deep.equal([]);  
    });
    
  });
});
