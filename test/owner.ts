import { expect } from "chai";
import { Address, Contract, Signer } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateKeyPair } from 'crypto';

let ownerContract: Contract<FactorySource["Owner"]>;
// let signer: Signer;
let publicKey: string;

describe.skip("Owner contract", async function () {
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
      const sampleData = await locklift.factory.getContractArtifacts("Owner");
      expect(sampleData.code).not.to.equal(undefined, "Code should be available");
      expect(sampleData.abi).not.to.equal(undefined, "ABI should be available");
      expect(sampleData.tvc).not.to.equal(undefined, "tvc should be available");
    });

    it("Deploy contract", async function () {
      const { contract } = await locklift.factory.deployContract({
        contract: "Owner",
        publicKey: publicKey,
        initParams: {},
        constructorParams: {
            elector: "0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a",
        //     ipPort: "123.0.123.0:5865",
        //     contactInfo: "test-owner",
        },
        value: locklift.utils.toNano(2)
      });
      ownerContract = contract;

      // expect(await locklift.provider.getBalance(ownerContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Get for owner", async function () {
      const response = await ownerContract.methods.get({}).call();
      expect(response.elector.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a', "Wrong is set");
    });
    
    it("Get elector for owner", async function () {
      const response = await ownerContract.methods.getElector({}).call();
      expect(response.value0.toString()).to.be.equal('0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a', "Wrong is set getElector");
    });
    
  });
});
