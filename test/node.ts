import { expect } from "chai";
import { Address, Contract, Signer } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateKeyPair } from 'crypto';

let nodeContract: Contract<FactorySource["Node"]>;
// let signer: Signer;
let publicKey: string;
let state: any;

describe("Node contract", async function () {
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
      const sampleData = await locklift.factory.getContractArtifacts("Node");
      expect(sampleData.code).not.to.equal(undefined, "Code should be available");
      expect(sampleData.abi).not.to.equal(undefined, "ABI should be available");
      expect(sampleData.tvc).not.to.equal(undefined, "tvc should be available");
    });

    it("Deploy contract", async function () {
      const { contract } = await locklift.factory.deployContract({
        contract: "Node",
        publicKey: publicKey,
        initParams: {},
        constructorParams: {
            elector: "0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a",
            ipPort: "123.0.123.0:5865",
            contactInfo: "test-node",
        },
        value: locklift.utils.toNano(2)
      });
      nodeContract = contract;

      // expect(await locklift.provider.getBalance(nodeContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Get node for node", async function () {
      const response = await nodeContract.methods.get({}).call();
      expect(response.elector.toString()).to.be.equal("0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a", "Wrong node Elector!");
      expect(response.ipPort.toString()).to.be.equal("123.0.123.0:5865", "Wrong node ipPort");
      expect(response.contactInfo.toString()).to.be.equal("test-node", "Wrong node contactInfo");
    });
    
    it("Get Elector for node", async function () {
      const response = await nodeContract.methods.getElector({}).call();
      expect(response.value0.toString()).to.be.equal("0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a", "Wrong node Elector!");
    });
    
    it("Get getIpPort for node", async function () {
      const response = await nodeContract.methods.getIpPort({}).call();
      expect(response.value0.toString()).to.be.equal("123.0.123.0:5865", "Wrong node ipPort!");
    });
    
    it("Get getContactInfo for node", async function () {
      const response = await nodeContract.methods.getContactInfo({}).call();
      expect(response.value0.toString()).to.be.equal("test-node", "Wrong node ContactInfo!");
    });

    it("Set and get ipPort for node", async function () {
      const newIpPort = "91.0.91.0:1234";
      await nodeContract.methods.
      setIpPort({ value: newIpPort }).
        sendExternal({ publicKey: publicKey });
      const response = await nodeContract.methods.getIpPort({}).call();
      expect(response.value0.toString()).to.be.equal(newIpPort.toString(), "Wrong ipPort is set");
    });

    it("Set and get ContactInfo for node", async function () {
      const newContactInfo = "Automation-test-node";
      await nodeContract.methods.
      setContactInfo({ value: newContactInfo }).
        sendExternal({ publicKey: publicKey });
      const response = await nodeContract.methods.getContactInfo({}).call();
      expect(response.value0.toString()).to.be.equal(newContactInfo.toString(), "Wrong ContactInfo is set");
    });
  });
});
