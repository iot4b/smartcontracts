pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;

contract Elector {
    string private _contractVersion = "v0.0.1";

    // список текущих нод на этом электоре
    address[] public _nodesCurrent; // List of current nodes
    address[] public _nodesNext; // List of next nodes
    address[] public _nodesParticipants ; // List of all nodes that want to participate in the election
    address[] public _nodesFault; // list of nodes who was fault in current round


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
        _nodesCurrent = defaultNodes;
    }

    // Устанавливаем список нод для текущего цикла
    function setNodes(address[] nodes) public onlyAccountOwner {
        _nodesCurrent = nodes;
    }

    // You can read from a state variable without sending a transaction.
    function get() public alwaysAccept view returns (
        address[] nodes
    ) {
        return (
            _nodesCurrent
        );
    }

    // todo возвращать версию текущего контракта
    function v() public alwaysAccept view returns (string contractVersion) {
        return _contractVersion;
    }


    /*
    func CurrentList() // список текущих нод на этом электоре
    func NextList()  // список нод на следующем электоре
    func Participants() // список всех нод желающих участвовать в электоре
    func TakeNextRound(address) //принять участие в следующем раунде
    func Election() // провести выборы для следующего раунда

    func ReportFaultNode(address) // сообщить о некорректной работе ноды
    func ProcessPaymentsNode(address) // обработать выплаты ноде
    func ProcessPaymentsVendor(address) // обработать выплаты вендору
    func Deposit() // пополнить депозит для девайса
    func WithdrawDeposit() // вывести депозит для девайса
    func 


    
    */
}