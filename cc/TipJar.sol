pragma solidity ^0.4.19;

contract TipJar {
    address owner;

    modifier ownerOnly {
        require(owner == msg.sender);
        _;   // <--- note the '_', which represents the modified function's body
    }

    function TipJar() public {  // contract's constructor function
        owner = msg.sender;
    }

    function changeOwner(address newOwner) public ownerOnly {
        owner = newOwner;
    }

    function withdraw() public ownerOnly {
        msg.sender.transfer(address(this).balance);
    }

    function deposit(uint256 amount) payable public {
        require(msg.value == amount);
    }

    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
}