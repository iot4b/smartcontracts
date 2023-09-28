pragma ever -solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {
    address public _node;
    address public _elector;
    address public _vendor;
    address[] public _owners;

    bool public _active;
    bool public _lock;
    bool  public _stat;

    string public _dtype;
    string public _version;
    string public _vendorName;
    string public _vendorData;

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    modifier onlyElectorContract() {
        require(msg.sender == _elector, 102);
        tvm.accept();
        _;
    }

    constructor(
        address node
    ) {
        tvm.accept();
        // setup addresses data
        _node = node;
    }

    // Set current node address for device
    function setNode(address value) public onlyElectorContract {
        _node = value;
    }

    // Get current node address for device
    function getNode() public alwaysAccept view returns (address) {
        return _node;
    }

    // Get current node address for device
    function getElector() public alwaysAccept view returns (address) {
        return _elector;
    }

    // Get current node address for device
    function getVendor() public alwaysAccept view returns (address) {
        return _vendor;
    }

    // Get current node address for device
    function getOwners() public alwaysAccept view returns (address[]) {
        return _owners;
    }
}
