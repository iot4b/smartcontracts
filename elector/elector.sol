pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Elector {
    // State variable to store a list of nodes
    string public nodeList;

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    // You need to send a transaction to write to a state variable.
    function set(string list) public alwaysAccept {
        nodeList = list;
    }

    // You can read from a state variable without sending a transaction.
    function get() public alwaysAccept view returns (string) {
        return nodeList;
    }
}