pragma solidity >=0.4.22 <0.6.0;

contract ERC223 {

    uint public totalSupply;
    function balanceOf(address who) constant returns (uint);
    function transfer(address to, uint value);
    function transfer(address to, uint value, bytes data);
    event Transfer(address indexed from, address indexed to, uint value, bytes data);
}

contract MyERC223 is ERC223 {
    constructor () public {
    }
}