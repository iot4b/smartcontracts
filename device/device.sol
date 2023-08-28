pragma solidity ^0.8.0;

interface IERC20 {
    function transfer(address recipient, uint256 amount) external returns (bool);
}

contract Device {
    string public dtype; // тип девайса
    string public status; // текущий статус девайса on/off
    string public locked; // девайс заблокирован
    string public version; // todo ???
    address public currentNode; // адрес контракта конкретной ноды
    address public vendorContract; // адрес контракта производителя девайса
    address public electorContract; // адрес контракта Elector'a
    address public nodeContract; // todo адрес ???
    address public ownerContract; // владелец девайса
    IERC20 public token;

    constructor(IERC20 _token) {
        ownerContract = msg.sender;
        token = _token;
    }

    // only-модификаторы определяют ограничения для выполнения функции
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

    // получить адрес текущей ноды
    function getNodeAddress() public view returns (address) {
        return currentNode;
    }

    // установить текущую ноду для девайса
    function setNodeAddress(address newNode) public onlyElectorContract {
        currentNode = newNode;
    }

    // получить тип девайса
    function getType() public view returns (string memory) {
        return dtype;
    }

    // todo получить версию прошивки девайса??? версии девайса???
    function getVersion() public view returns (string memory) {
        return version;
    }

    // указать версию прошивки. нода получает инфу от девайса и
    function setVersion(string newVersion) public onlyCurrentNode {
        version = newVersion;
    }

    // установить статус девайса.может устанавливать только текущая нода
    function setStatus(string memory newStatus) public onlyCurrentNode {
        status = newStatus;
    }

    // текущий статус девайса
    function getStatus() public view returns (string memory) {
        return status;
    }

    // списание за обслуживание. с баланса девайса списывается на баланс владельца ноды
    function pay(uint256 amount) public onlyCurrentNode {
        require(token.transfer(currentNode, amount), "Transfer failed.");
    }

    function checkLock() public view returns (bool) {
        return locked;
    }

    function setLock() public onlyCurrentNode {
        locked = true;
    }

    function unlock() public onlyCurrentNode {
        locked = false;
    }
}
