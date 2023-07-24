
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Device {
    // State variable to store a number
    string public node;

    // You need to send a transaction to write to a state variable.
    function set(string _node) public {
        node = _node;
    }

    // You can read from a state variable without sending a transaction.
    function get() public view returns (string) {
        return node;
    }
}