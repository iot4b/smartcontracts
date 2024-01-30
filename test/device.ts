import { expect } from "chai";
import { Address, Contract, Signer } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateKeyPair } from 'crypto';

let deviceContract: Contract<FactorySource["Device"]>;
// let signer: Signer;
let publicKey: string;
let state: any;

describe.skip("Device contract", async function () {
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
            owners: ["6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b"],
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
        setNode({ value: newNode }).
        sendExternal({ publicKey: publicKey });
      const response = await deviceContract.methods.getNode({}).call();
      
      expect(response.value0.toString()).to.be.equal(newNode.toString(), "Wrong node is set");
    });
    
    it("Get Elector for device", async function () {
      const response = await deviceContract.methods.getElector({}).call();
      expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a');
    });

    it("Get Vendor for device", async function () {
      const response = await deviceContract.methods.getVendor({}).call();
      console.log(response);
      expect(response.value0.toString()).to.be.equal('0:cf59bb48dac2b1234bce4b5c8108f8c884852ca1333065caa16adf4a86051337');
    });
    
    it("Get Owner for device", async function () {
      const response = await deviceContract.methods.getOwners({}).call();
      console.log(response);
      expect(response.value0.toString()).to.be.equal('6bbadda1506aeb790dcc8a03aa94c1b25f81edf20892c24cc81a062e788bfa7b');
    });
    
    // it("Get VendorData for device", async function () {
    //   const response = await deviceContract.methods.getVendorData({}).call();
    //   console.log(response);
    //   // expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a');
    // });

    it("Get data for device", async function () {
      const response = await deviceContract.methods.get({}).call();
      state = response;
      // expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a');
    });
    
    it("Get lock for device", async function () {
      const response = await deviceContract.methods.setLock({lock: true}).call();
      console.log(response);
      const checkState = await deviceContract.methods.get({}).call();
      console.log(checkState);
      // expect(checkState.lock).to.be.equal(true);
    });
    
    it("Get active for device", async function () {
      const response = await deviceContract.methods.setActive({active: true}).call();
      console.log(response);
      const checkState = await deviceContract.methods.get({}).call();
      console.log(checkState);
      // expect(checkState.active).to.be.equal(true);
    });
    
    it("Get stat for device", async function () {
      const response = await deviceContract.methods.setStat({stat: true}).call();
      console.log(response);
      const checkState = await deviceContract.methods.get({}).call();
      console.log(checkState);
      // expect(checkState.stat).to.be.equal(true);
    });
  });
});
