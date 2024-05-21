import { expect } from "chai";
import { Address, Contract } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateSignKeys } from "../scripts/util";

let deviceContract: Contract<FactorySource["Device"]>;
let publicKey: string;

describe("Device contract", async function () {
  before(async () => {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);
    publicKey = signer.publicKey;
  });
  describe("Contracts", async function () {
    it("Load contract factory", async function () {
      const sampleData = await locklift.factory.getContractArtifacts("Device");
      expect(sampleData.code).not.to.equal(undefined, "Code should be available");
      expect(sampleData.abi).not.to.equal(undefined, "ABI should be available");
      expect(sampleData.tvc).not.to.equal(undefined, "tvc should be available");
    });

    it("Deploy contract", async function () {
      const { contract } = await locklift.factory.deployContract({
        contract: "Device",
        publicKey: publicKey,
        initParams: {},
        constructorParams: {
            elector: new Address("0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a"),
            vendor: new Address("0:cf59bb48dac2b1234bce4b5c8108f8c884852ca1333065caa16adf4a86051337"),
            owners: [
                [
                  "0x6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b", // public key
                  "0:0000000000000000000000000000000000000000000000000000000000000000" // address, zeroes if not provided
                ]
            ],
            dtype: "test-device",
            version: "0.1",
            vendorName: "Apple",
            vendorData: "{\"serialNumber\":\"DSF34-G4FG34G\"}"
        },
        value: locklift.utils.toNano(2)
      });
      deviceContract = contract;

      // expect(await locklift.provider.getBalance(deviceContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Set and get node for device", async function () {
      const newNode = new Address("0:675a6d63f27e3f24d41d286043a9286b2e3eb6b84fa4c3308cc2833ef6f54d68"); // new node address
      await deviceContract.methods.
        setNode({ node: newNode }).
        sendExternal({ publicKey: publicKey });
      const response = await deviceContract.methods.getNode({}).call();
      expect(response.value0.toString()).to.be.equal(newNode.toString(), "Wrong node is set");
    });

    it("Get Elector for device", async function () {
      const response = await deviceContract.methods.getElector({}).call();
      expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a'); // TO DO - fix to initual data type
    });

    it("Get Vendor for device", async function () {
      const response = await deviceContract.methods.getVendor({}).call();
      expect(response.value0.toString()).to.be.equal('0:cf59bb48dac2b1234bce4b5c8108f8c884852ca1333065caa16adf4a86051337'); // TO DO - fix to initual data type
    });

    it("Get Owner for device", async function () {
      const response = await deviceContract.methods.getOwners({}).call();
      expect(BigInt(response.value0[0][0]).toString(16)).to.be.equal('6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b'); // TO DO - fix to initual data type
    });

    // it("Get VendorData for device", async function () {
    //   const response = await deviceContract.methods.getVendorData({}).call();
    //   console.log(response);
    //   // expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a');
    // });

    it("Get data for device", async function () {
      const response = await deviceContract.methods.get({}).call();
      expect(response.node.toString()).to.be.equal('0:675a6d63f27e3f24d41d286043a9286b2e3eb6b84fa4c3308cc2833ef6f54d68');
      expect(response.elector.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a');
      expect(response.vendor.toString()).to.be.equal('0:cf59bb48dac2b1234bce4b5c8108f8c884852ca1333065caa16adf4a86051337');
      expect(BigInt(response.owners[0][0]).toString(16)).to.be.equal('6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b');
      expect(response.dtype).to.be.equal('test-device');
      expect(response.version).to.be.equal('0.1');
      expect(response.vendorName).to.be.equal('Apple');
    });

    it("Get lock for device", async function () {
      await deviceContract.methods.setLock({lock: true}).sendExternal({ publicKey: publicKey });
      const checkState = await deviceContract.methods.get({}).call();
      expect(checkState.lock).to.be.equal(true);
    });

    it("Get active for device", async function () {
      await deviceContract.methods.setActive({active: true}).sendExternal({ publicKey: publicKey });
      const checkState = await deviceContract.methods.get({}).call();
      setTimeout(() => console.log('Привет'), 10000);
      expect(checkState.active).to.be.equal(true);
    });

    it("Get stat for device", async function () {
      await deviceContract.methods.setStat({stat: true}).sendExternal({ publicKey: publicKey });
      const checkState = await deviceContract.methods.get({}).call();
      console.log(checkState);
      expect(checkState.stat).to.be.equal(true);
    });
  });
});
