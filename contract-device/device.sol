pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {

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
