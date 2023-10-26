pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Elector {
    string private _contractVersion = "v0.0.1";

    // список текущих нод на этом электоре
    address[] public _nodes; // List of current nodes

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    // Only contract owner
    modifier onlyAccountOwner {
//        msg.sender == address();
        tvm.accept();
        _;
    }

    // передаем ноды по умолчанию
    constructor(
        address[] defaultNodes
    ) {
        tvm.accept();
        _nodes = defaultNodes;
    }

    // Устанавливаем список нод для текущего цикла
    function setNodes(address[] nodes) public onlyAccountOwner {
        _nodes = nodes;
    }

    // You can read from a state variable without sending a transaction.
    function get() public alwaysAccept view returns (
        address[] nodes
    ) {
        return (
            _nodes
        );
    }

    // todo возвращать версию текущего контракта
    function v() public alwaysAccept view returns (string contractVersion) {
        return _contractVersion;
    }
}