pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Node {
    string public location;
    string public ipPort;

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    function setLocation(string _location) public {
        location = _location;
    }

    function setIpPort(string _ipPort) public {
        ipPort = _ipPort;
    }
}