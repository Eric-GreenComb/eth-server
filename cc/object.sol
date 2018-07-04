pragma solidity ^0.4.11;

contract Object {

    string public value;

    string public key;

    function ModifyValue(string val) public{
        value = val;
    }
}