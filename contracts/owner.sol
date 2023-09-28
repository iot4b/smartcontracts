pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Owner {
    string private _contractVersion = "v0.0.1";

    // адрес электора
    address public _elector;

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
        address elector
    ) {
        tvm.accept();
        _elector = elector;
    }

    function get() public alwaysAccept view returns (
        address elector
    ) {
        return (
            _elector
        );
    }

    function getElector() public alwaysAccept view returns (address) {
        return _elector;
    }

    // todo возвращать версию текущего контракта
    function v() public alwaysAccept view returns (string contractVersion) {
        return _contractVersion;
    }
}