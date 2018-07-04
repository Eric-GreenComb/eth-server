pragma solidity ^0.4.19;

contract ObjectInterface {
    function ModifyValue(string val) public;
}

contract CallObject {

    function ModifyValue(address addrObj,string val) public{

        ObjectInterface(addrObj).ModifyValue(val);
    }
}