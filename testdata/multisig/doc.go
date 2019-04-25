// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package multisig

/*
https://medium.com/@yenthanh/list-of-multisig-wallet-smart-contracts-on-ethereum-3824d528b95e

List of MultiSig Wallet Smart contracts

Here I will list some implementation for MultiSig Wallet that are verified and widely used. Please keep in mind that you shouldn’t change the implementation as the code is tested carefully to prevent any problem unless you really know what you are doing. Of course I can not guarantee that those smart contracts are 100% bug free as it is made by human anyway :) So always test everything carefully before deploy it to the mainnet.
And here is the list:

ConsenSys’ Multisig Wallet:
This is one of the simplest implementation for the Multisig Wallet, the code is easy to understand and use, basically you can use the multisig wallet to do anything an individual account can do (except deploying contract, of course!). It is widely used and one of the biggest wallet deployment has 200k ETHER! (reviewed on the time writing this post). Its drawback is lacking of test and not updated for a long time, the code still use solidity ver 0.4.10!
Gnosis’ Multisig Wallet:
This is an upgrade implementation for the ConsenSys’ Multisig Wallet. The wallet still maintain the simple logic and code as its a.The code is structured as a Truffle project with tests and active developments, this is a good alternative choice if you don’t want to use the Wallet from ConsenSys.
BitGo’s Multisig Wallet:
This implementation is also a good choice. The project is implemented with Truffle, tested with an active community. The wallet has more complex logic than the Gnosis’s and ConsenSys’s but still easy for you to understand and logic. One of the biggest advantage of this wallet is ERC20-Token Compatibility, owner can use defined function to transfer Token easier than the other wallets (yes, of course we can use other wallet to transfer token or do any advanced stuffs). Notice that this Multisig Wallet is implemented using 2-of-3 signing configuration, meaning the wallet has exactly 3 owners and require 2 agreements to proceed the transaction.
Ethereum Dapp’s Multisig Wallet (Sourcecode on Etherscan):
If you use Ethereum Wallet or Mist, you can see that the App also provide you a method to deploy a Multisign Wallet. The interesting part is this wallet is compatible with the Ethereum Wallet App and Mist App, which means you can call the function send transaction or confirm easily. Although there are no documents about how to use this wallet, we can analyze its implementation here (a little bit different from the one committed on ethereum’s repo) and understand the logic clearly. There are some function has a special implementation that requires tricky way to proceed, such as ‘changeRequirements’ needs to be call by many owners and mined in the same block to be proceeded.
Parity’s Multisig Wallet (NOT RECOMMENDED):
Parity UI App also has a similar feature as Ethereum Wallet which helps us deploying the Multisign Wallet easily. Sad news is in Nov 6th, 2017 the parity’s multisig wallet is hacked which make $300mil frozen forever, the problem is because of all the deployed multisign wallet use a library implemented by Parity which isn’t initialized correctly. A person luckily become the owner of the library because of this bug and try to stole the money by calling the ‘kill’ function, he couldn’t get the money but make all smart contracts used this library died. That’s why all the implementation for Multisig Wallet I listed here has no library dependencies in the code to prevent this scenario happen again. I don’t know if the Parity has changed the way they deploy the smart contracts for MultiSig Wallet or not but I would not recommend people to use the Parity to generate the Multisig Wallet now (anyway, Parity is still a good choice for Ethereum Node now).

*/
