pragma solidity ^0.8.0;

contract Device {
    address public currentNode;
    string public dtype;
    address public vendorContract;
    address public nodeContract;
    address public ownerContract;
    string public status;

    constructor() {
        ownerContract = msg.sender;
    }

    modifier onlyNodeContract() {
        require(msg.sender == nodeContract, "Only the nodeContract can call this method.");
        _;
    }

    function getNode() public view returns (address) {
        return currentNode;
    }

    function setNode(address newNode) public onlyNodeContract {
        currentNode = newNode;
    }

    function getType() public view returns (string memory) {
        return dtype;
    }

    function setStatus(string memory newStatus) public {
        status = newStatus;
    }

    function getStatus() public view returns (string memory) {
        return status;
    }
}
