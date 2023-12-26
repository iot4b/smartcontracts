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
      const { deviceContract } = await locklift.factory.deployContract({
        contract: "Device",
        publicKey: signer.publicKey,
        initParams: {},
        constructorParams: {
            elector: "0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a",
            vendor: "0:cf59bb48dac2b1234bce4b5c8108f8c884852ca1333065caa16adf4a86051337",
            owners: ["6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b"],
            dtype: "test-device",
            version: "0.1",
            vendorName: "Apple",
            vendorData: "{\"serialNumber\":\"DSF34-G4FG34G\"}"
        },
        value: locklift.utils.toNano(2)
      });

      expect(await locklift.provider.getBalance(deviceContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Interact with contract", async function () {
      const newNode = new Address("0:675a6d63f27e3f24d41d286043a9286b2e3eb6b84fa4c3308cc2833ef6f54d68"); // new node address

      await deviceContract.methods.
      setNode({ value: newNode }).
      sendExternal({ publicKey: signer.publicKey });

      const response = await deviceContract.methods.getNode({}).call();

      expect(response.value0).to.be.equal(newNode, "Wrong node state");
    });
  });
});
