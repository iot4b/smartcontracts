pragma solidity ^0.8.0;

interface IERC20 {
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
}

contract Device {
    string public dtype;
    string public status;
    string public version;
    address public currentNode;
    address public vendorContract;
    address public electorContract;
    address public nodeContract;
    address public ownerContract;
    IERC20 public token;

    constructor(IERC20 _token) {
        ownerContract = msg.sender;
        token = _token;
    }

    modifier onlyNodeContract() {
        require(msg.sender == nodeContract, "Only the nodeContract can call this method.");
        _;
    }

    modifier onlyCurrentNode() {
        require(msg.sender == currentNode, "Only the currentNode can call this method.");
        _;
    }

    modifier onlyElectorContract() {
        require(msg.sender == electorContract, "Only the electorContract can call this method.");
        _;
    }

    function getNode() public view returns (address) {
        return currentNode;
    }

    function setNode(address newNode) public onlyElectorContract {
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

    function getPayment(uint256 amount) public onlyCurrentNode {
        require(token.transferFrom(ownerContract, currentNode, amount), "Transfer failed.");
    }
}
