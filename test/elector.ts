import { expect } from "chai";
import { Address, Contract } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateSignKeys } from "../scripts/util";

let electorContract: Contract<FactorySource["Elector"]>;
let publicKey: string;

function timeout(fn, delay) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      try {
        fn();
        resolve();
      } catch(err) {
        reject(err);
      } 
    }, delay);
  });
}

describe("Elector contract", async function () {
  before(async () => {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);
    publicKey = signer.publicKey;
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

    it("Set nodes for elector", async function () {
      await electorContract.methods.setNodes({
        _nodes: [
            [new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'), true],
            [new Address('0:86429800dd5b8ddc9a1283341b106cdb7acb2807c4e5f91e523c2803e6c76ddd'), true],
            [new Address('0:e986b8305e5d46cc221cc9e14785bfe361b8558104396bdc082fa4c6321ffc68'), true]
        ]}).sendExternal({ publicKey: publicKey });
        const responseCheck = await electorContract.methods.get().call();
        expect(responseCheck.nodesCurrent).not.to.deep.equal([]);  
    });

    it("Set participantList for elector", async function () {
      await electorContract.methods.takeNextRound({
          _address: new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'),
        }).sendExternal({ publicKey: publicKey });
      const responseCheck = await electorContract.methods.get().call();
      expect(responseCheck.nodesParticipants).not.to.deep.equal([]);  
    });
    
    it.skip("Election test of elector", async function () {
      
      await electorContract.methods.takeNextRound({
          _address: new Address('0:4a2158bd934f0f199224b89dd58f8b20ad73a160ef06ca67d55a63fc8d4b0a26'),
        }).sendExternal({ publicKey: publicKey });
      // await electorContract.methods.election().call();
      const responseCheck = await electorContract.methods.election().call();
      console.log(responseCheck);
      const responseAfter = await electorContract.methods.get().call();
      // expect(responseAfter.nodesParticipants).not.to.deep.equal([]);
      console.log(responseAfter.nodesNext);
      // expect(responseAfter.nodesNext).
    });
    
  });
});
