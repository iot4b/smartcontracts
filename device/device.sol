pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Device {
    string  public dtype; // тип девайса
    string  public status; // текущий статус девайса on/off
    bool    public locked; // девайс заблокирован
    string  public version; // todo ???
    address public currentNode; // адрес контракта конкретной ноды
    address public vendorContract; // адрес контракта производителя девайса
    address public electorContract; // адрес контракта Elector'a
    address public nodeContract; // todo адрес ???
    address public ownerContract; // владелец девайса

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        tvm.accept();
        _;
    }

    modifier onlyCurrentNode() {
//        require(msg.sender == currentNode, "Only the currentNode can call this method.");
        tvm.accept();
        _;
    }

    // получить адрес текущей ноды
    function getNodeAddress() public alwaysAccept view returns (address) {
        return currentNode;
    }

    // установить текущую ноду для девайса
    function setNodeAddress(address newNode) public alwaysAccept {
        currentNode = newNode;
    }

    function getType() public alwaysAccept view returns (string memory) {
        return dtype;
    }

    // todo получить версию прошивки девайса??? версии девайса???
    function getVersion() public alwaysAccept view returns (string memory) {
        return version;
    }

    function setStatus(string newStatus) public onlyCurrentNode {
        status = newStatus;
    }

    function getStatus() public view returns (string memory) {
        return status;
    }

    // списание за обслуживание. с баланса девайса списывается на баланс владельца ноды
    function pay(uint256 amount) public onlyCurrentNode {
//        require(token.transfer(currentNode, amount), "Transfer failed.");
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
