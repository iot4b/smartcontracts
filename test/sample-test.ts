import { expect } from "chai";
import { Address, Contract, Signer } from "locklift";
import { FactorySource } from "../build/factorySource";

let deviceContract: Contract<FactorySource["Device"]>;
let signer: Signer;

describe("Test Sample contract", async function () {
  before(async () => {
    signer = (await locklift.keystore.getSigner("0"))!;
  });
  describe("Contracts", async function () {
    it("Load contract factory", async function () {
      const sampleData = await locklift.factory.getContractArtifacts("Device");

      expect(sampleData.code).not.to.equal(undefined, "Code should be available");
      expect(sampleData.abi).not.to.equal(undefined, "ABI should be available");
      expect(sampleData.tvc).not.to.equal(undefined, "tvc should be available");
    });

    it("Deploy contract", async function () {
      const INIT_STATE = 0;
      const { contract } = await locklift.factory.deployContract({
        contract: "Device",
        publicKey: signer.publicKey,
        initParams: {
          _nonce: locklift.utils.getRandomNonce(),
        },
        constructorParams: {
            version: '0,1',
        },
        value: locklift.utils.toNano(2),
      });
      deviceContract = contract;

      expect(await locklift.provider.getBalance(deviceContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Interact with contract", async function () {
      const newNode = ; // new node address
 
      await deviceContract.methods.setNode().sendExternal({ publicKey: signer.publicKey });

      const response = await deviceContract.methods.getNode({}).call();

      expect(response).to.be.equal(newNode, "Wrong node state");
    });
  });
});
