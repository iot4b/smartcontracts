pragma ever-solidity ^0.71.0;
pragma AbiHeader expire;
//pragma AbiHeader pubkey;

contract Vendor {
  string private _contractVersion = "0.0.1";

  address public _elector;
  // название производителя
  string public _vendorName;
  // доля вендора и нод от прибыли. max 100, min 0 %, default 50
  uint public _profitShare;
  // контактные данные производителя
  string public _contactInfo;

  // Modifier that allows public function to accept all external calls.
  modifier alwaysAccept {
    tvm.accept();
    _;
  }

  /// @dev Contract constructor.
  constructor(
    address elector,
    string vendorName,
    uint profitShare,
    string contactInfo
  ) {
    // todo check pubkey
    // check that contract's public key is set
//        require(tvm.pubkey() != 0, 101);
    // Check that message has signature (msg.pubkey() is not zero) and message is signed with the owner's private key
//        require(msg.pubkey() == tvm.pubkey(), 102);
    // check uint value
//    require(profitShare >= 0, "profitShare must be in [0; 100]");
//    require(profitShare <= 100, "profitShare must ben [0; 100]");
    require(profitShare >= 0, 102);
    require(profitShare <= 100, 102);
    tvm.accept();

    // set initial data
    _elector = elector;
    _vendorName = vendorName;
    _profitShare = profitShare;
    _contactInfo = contactInfo;
  }

  function get() public alwaysAccept view returns (
    address elector,
    string vendorName,
    uint profitShare,
    string contactInfo
  ) {
    return (
      _elector,
      _vendorName,
      _profitShare,
      _contactInfo
    );
  }

  function getElector() public view returns (address) {
    return _elector;
  }

  function getVendorName() public view returns (string) {
    return _vendorName;
  }

  function getProfitShare() public view returns (uint) {
    return _profitShare;
  }

  function getContactInfo() public view returns (string) {
    return _contactInfo;
  }

  function setVendorName(string value) public {
    _vendorName = value;
  }

  function setProfitShare(uint value) public {
    require(value >= 0, 102);
    require(value <= 100, 102);
    tvm.accept();

    _profitShare = value;
  }

  function setContactInfo(string value) public {
    _contactInfo = value;
  }

  // todo возвращать версию текущего контракта
  function v() public alwaysAccept view returns (string contractVersion) {
    return _contractVersion;
  }
}