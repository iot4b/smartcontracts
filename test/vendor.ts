import { expect } from "chai";
import { Contract } from "locklift";
import { FactorySource } from "../build/factorySource";
import { generateSignKeys } from "../scripts/util";

let vendorContract: Contract<FactorySource["Vendor"]>;
let publicKey: string;

describe("Vendor contract", async function () {
  before(async () => {
    const signer = await generateSignKeys();
    locklift.keystore.addKeyPair(signer);
    publicKey = signer.publicKey;
  });
  describe("Contracts", async function () {
    it("Load contract factory", async function () {
      const sampleData = await locklift.factory.getContractArtifacts("Vendor");
      expect(sampleData.code).not.to.equal(undefined, "Code should be available");
      expect(sampleData.abi).not.to.equal(undefined, "ABI should be available");
      expect(sampleData.tvc).not.to.equal(undefined, "tvc should be available");
    });

    it("Deploy contract", async function () {
      // Failed to wait for next block for 0:ece57bcc6c530283becbbd8a3b24d3c5987cdddc3c8b7b33be6e4a6312490415
      const { contract } = await locklift.factory.deployContract({
        contract: "Vendor",
        publicKey: publicKey,
        initParams: {},
        constructorParams: {
            elector: "0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a",
            vendorName: "testVendor-Automation",
            contactInfo: "testContact-vendor",
            profitShare: 100,
        },
        value: locklift.utils.toNano(2)
      });
      vendorContract = contract;

      // expect(await locklift.provider.getBalance(vendorContract.address).then(balance => Number(balance))).to.be.above(0);
    });

    it("Get vendor for vendor", async function () {
      const response = await vendorContract.methods.get({}).call();
      expect(response.elector.toString()).to.be.equal("0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a", "Wrong vendor Elector!");
      expect(response.vendorName.toString()).to.be.equal("testVendor-Automation", "Wrong vendor vendorName");
      expect(response.contactInfo.toString()).to.be.equal("testContact-vendor", "Wrong vendor contactInfo");
      expect(response.profitShare).to.be.equal('100', "Wrong vendor profitShare");
    });
    
    it("Get Elector for vendor", async function () {
      const response = await vendorContract.methods.get({}).call();
      expect(response.elector.toString()).to.be.equal("0:da995a0f7e2f75457031cbc016d7cba6fc65b617a94331eb54c349af15e95d1a", "Wrong vendor Elector!");
    });
    
    it("Get Name for vendor", async function () {
      const response = await vendorContract.methods.get({}).call();
      expect(response.vendorName.toString()).to.be.equal("testVendor-Automation", "Wrong vendor name!");
    });
    
    it("Get getProfitShare for vendor", async function () {
      const response = await vendorContract.methods.get({}).call();
      expect(response.profitShare.toString()).to.be.equal("100", "Wrong vendor ProfitShare!");
    });
    
    it("Get getContactInfo for vendor", async function () {
      const response = await vendorContract.methods.get({}).call();
      expect(response.contactInfo.toString()).to.be.equal("testContact-vendor", "Wrong vendor ProfitShare!");
    });

    it("Set and get Vendor name for vendor", async function () {
      const newValue = "BRZ";
      await vendorContract.methods.
      setVendorName({ value: newValue }).
        sendExternal({ publicKey: publicKey });
      const response = await vendorContract.methods.get({}).call();
      expect(response.vendorName.toString()).to.be.equal(newValue.toString(), "Wrong Vendor name is set");
    });

    it("Set and get ProfitShare for vendor", async function () {
      const newValue = "50";
      await vendorContract.methods.
      setProfitShare({ value: newValue }).
        sendExternal({ publicKey: publicKey });
      const response = await vendorContract.methods.get({}).call();
      expect(response.profitShare.toString()).to.be.equal(newValue.toString(), "Wrong ProfitShare is set");
    });

    it("Set and get ContactInfo for vendor", async function () {
      const newValue = "This is bew contact info";
      await vendorContract.methods.
      setContactInfo({ value: newValue }).
        sendExternal({ publicKey: publicKey });
      const response = await vendorContract.methods.get({}).call();
      expect(response.contactInfo.toString()).to.be.equal(newValue.toString(), "Wrong ContactInfo is set");
    });

    it("Negative case: Set and get ProfitShare for vendor for >100", async function () {
      locklift.tracing.setAllowedCodes({ compute: [102], action: [102] });
      const newValue = "150";
      await vendorContract.methods.
      setProfitShare({ value: newValue }).
        sendExternal({ publicKey: publicKey });
      const response = await vendorContract.methods.get({}).call();
      expect(response.profitShare.toString()).not.to.be.equal(newValue.toString(), "Wrong ProfitShare is set");
    });
  });
});
