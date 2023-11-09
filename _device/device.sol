pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {
    string private _contractVersion = "v1.0.1";

    address public _node;
    address public _elector;
    address public _vendor;

    string[] public _owners; // public keys of device owners

    bool public _active;
    bool public _lock;
    bool public _stat;

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

    modifier onlyNodeContract() {
        require(msg.sender == _node, 102);
        tvm.accept();
        _;
    }

    modifier onlyVendorContract() {
        require(msg.sender == _vendor, 102);
        tvm.accept();
        _;
    }

    constructor(
        address elector,
        address vendor,

        string[] owners,

        string dtype,
        string version,
        string vendorName,
        string vendorData
    ) {
        tvm.accept();
        // setup addresses data
        _elector = elector;
        _vendor = vendor;

        _owners = owners;

        _active = false;
        _lock = false;
        _stat = false;

        _dtype = dtype;
        _version = version;
        _vendorName = vendorName;
        _vendorData = vendorData;
    }

    // get all contract data
    function get() public alwaysAccept view returns (
        address node,
        address elector,
        address vendor,
        string[] owners,
        bool active,
        bool lock,
        bool stat,
        string dtype,
        string version,
        string vendorName,
        string vendorData
    ) {
        return (
            _node,
            _elector,
            _vendor,
            _owners,
            _active,
            _lock,
            _stat,
            _dtype,
            _version,
            _vendorName,
            _vendorData
        );
    }

    // Set current node address for device
    function setNode(address value) public alwaysAccept {
        _node = value;
    }

    // Get current node address for device
    function getNode() public alwaysAccept view returns (address) {
        return _node;
    }

    // Get elector address for device
    function getElector() public alwaysAccept view returns (address) {
        return _elector;
    }

    // Get vendor address for device
    function getVendor() public alwaysAccept view returns (address) {
        return _vendor;
    }

    function getVendorData() public onlyVendorContract view returns (address) {
        return _vendorData;
    }

    // Get public keys of device owners
    function getOwners() public alwaysAccept view returns (string[]) {
        return _owners;
    }

    // Get contract version
    function v() public alwaysAccept view returns (string contractVersion) {
        return _contractVersion;
    }
}
