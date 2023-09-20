pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {
    string public node;
    string public elector;
    string public vendor;
    // todo как в солидити обозначаются массивы?
    address  public owners;
    bool public active;
    bool public lock;
    bool  public stat;
    string public Type;
    string public version;
    string public vendorName;
    string public vendorData;


    string public node; // Current node of the device

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    // Set current node address for device
    function setNode(string value) public alwaysAccept {
        node = value;
    }

    // Get current node address for device
    function getNode() public alwaysAccept view returns (string) {
        return node;
    }
}
