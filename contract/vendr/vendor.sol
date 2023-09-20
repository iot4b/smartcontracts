// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Vendor {
    string  public vendorName;
    string  public contactInfo;
    int     public profitShare;

    /// @dev Contract constructor.
    constructor() {
        // check that contract's public key is set
        require(tvm.pubkey() != 0, 101);
        // Check that message has signature (msg.pubkey() is not zero) and message is signed with the owner's private key
        require(msg.pubkey() == tvm.pubkey(), 102);
        tvm.accept();
    }

    // You need to send a transaction to write to a state variable.
    function set(string _node) public {
        node = _node;
    }

    // You can read from a state variable without sending a transaction.
    function get() public view returns (string) {
        return node;
    }
}