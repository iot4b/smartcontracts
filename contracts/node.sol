// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract Node {
    string public location;
    string public ipPort;

    function setLocation(string _location) public {
        location = _location;
    }

    function setIpPort(string _location) public {
        location = _location;
    }

    function getLocation() public view returns (string) {
        return location;
    }

    function getIpPort() public view returns (string) {
        return ipPort;
    }
}