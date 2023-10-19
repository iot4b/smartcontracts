pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Node {
    string private _contractVersion = "v0.0.1";

    address public _elector;
    // geo position
    string public _location;
    // ip:port
    string public _ipPort;
    // owner contact info
    string public _contactInfo;

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    // Modifier that allows public function to accept only Elector calls.
    modifier onlyElector {
        require(msg.sender == _elector, 102);
        tvm.accept();
        _;
    }

    constructor(
        address elector,
        string location,
        string ipPort,
        string contactInfo
    ) {
        tvm.accept();

        _elector = elector;
        _location = location;
        _ipPort = ipPort;
        _contactInfo = contactInfo;
    }

    function get() public alwaysAccept view returns (
        address elector,
        string location,
        string ipPort,
        string contactInfo
    )  {
        return (
            _elector,
            _location,
            _ipPort,
            _contactInfo
        );
    }

    function getElector() public alwaysAccept view returns (address) {
        return _elector;
    }

    function getLocation() public alwaysAccept view returns (string) {
        return _location;
    }

    function getIpPort() public alwaysAccept view returns (string) {
        return _ipPort;
    }

    function getContactInfo() public alwaysAccept view returns (string) {
        return _contactInfo;
    }

    // todo setter only account owner can use
    function setLocation(string value) public alwaysAccept {
        _location = value;
    }

    function setIpPort(string value) public alwaysAccept {
        _ipPort = value;
    }

    function setContactInfo(string value) public alwaysAccept {
        _contactInfo = value;
    }

    // todo возвращать версию текущего контракта
    function v() public alwaysAccept view returns (string contractVersion) {
        return _contractVersion;
    }
}