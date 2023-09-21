pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {
    address public node;
    address public elector;
    address public vendor;
    address[] public owners;

    bool public active;
    bool public lock;
    bool  public stat;

    string public dtype;
    string public version;
    string public vendorName;
    string public vendorData;

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    // Set current node address for device
    function setNode(address value) public alwaysAccept {
        node = value;
    }

    // Get current node address for device
    function getNode() public alwaysAccept view returns (address) {
        return node;
    }
}
