# Etherniti: High Performance Ethereum REST API

<p class="centered">
  <img src="./images/app.png" width="100px" />
</p>

## Executive summary

Etherniti is a Multitenant High Performance Ethereum and Quorum compatible REST API enabling fast prototyping features. Also compatible with QuikNode, Infura, Alastria, Geth and Parity.


## Why Etherniti?

We started this project long time ago because we saw a huge lack on current alternatives for DApps implementations. At the time we began, most tutorials and interactions ways were based on Web3.js (which is a really good library indeed), and Metamask (one of the best browser based wallets out there). However, we felt someting was missing. There were no easy ways to create a DApp that did not use Metamask for signing process or even web3 and their provider implementation for interaction. 

### Reason 1: Metamask lack of automatic interaction.

**MetaMask** is a bridge that allows you to visit the distributed web of tomorrow in your browser today. It allows you to run Ethereum dApps right in your browser without running a full Ethereum node.

This was the main reason to start the project. Metamask is a good starting point for developers, but still it lacks of automatic interaction from webapps. Developers cannot directly integrate a 'Buy' button that makes a transaction throw Ethereum network without having to confirm it via metamask. There are still ways to do this, but they are not as simple as running metamask. Thats wht, we still need proper tools to allow developers and the community to have an easy and reliable way to interact with decentralized Ethereum networks in a more automatic way.

### Reason 2: Web3js needs a provider for everything

Even this is a completely logic requirement (because you need to own an ethereum client like geth or parity in order to interact), for those developers that just want to get information from ethererum, read transactions, account balances, etc..there is no a real need to have a node/peer of its own (unless you need to have a highly trust about responses). 

With Etherniti, users will be able to query Ethereum and Quorum ledgers as any other online service like Twitter API, Twilio API, etc.

### Reason 3: Infura is not enough

**Infura** is a hosted Ethereum node cluster that lets your users run your application without requiring them to set up their own Ethereum node or wallet. You may not be familiar with Infura by name, but if you've used MetaMask then you've used Infura, as it is the Ethereum provider that powers MetaMask.

For those situations in where you need to have an easy and fast access to Ethereum, Infura is a good choice. 

**However** there are still some issue to tackle. Consensys Infura, adds an abstraction layer so that you dont need to own a ethereum node to interact. 
This is very closer to Etherniti goals, but still, **Infura only provides centralized support to public Ethereum networks leaving aside Quorum networks and private implementations of both technologies**.

<div
	class="alert alert-success"
	role="alert"
	style="text-align: justify;">
	We encourage you to read this post at Medium <a href="https://media.consensys.net/why-infura-is-the-secret-weapon-of-ethereum-infrastructure-af6fc7c77052" target="_blank" rel="noopener">https://media.consensys.net/why-infura-is-the-secret-weapon-of-ethereum-infrastructure-af6fc7c77052</a> and think about whether <b>Infura is centralizing Ethereum access</b> or not.
</div>

## Our vision: the **Etherniti Project**

Etherniti Project is a Multitenant High Performance Ethereum and Quorum compatible REST API.

Etherniti allows companies, developers and startups to interact with Ethereum in a very easy and friendly way. We add an abstraction layer that allows a fast, reliable and secure interaction with Ethereum networks and their smart contracts. We provide a fully featured High Performance REST API to operate ethereum as its usually done with other tools such as Web3js, Metamask, etc.

In other words, we provide the techology to have an easy, secure, reliable and high performance access to your custom private Ethereum or Quorum network, so you can focus on designing DApps rather than focusing on architecture and communication protocol of distributed ledgers.

## Usage scenarios

Etherniti can be very useful in several scenarios, such as:

* Have a consistent REST API style secure middleware for communicating with Ethereum and Quorum networks.
* Deploy our DApps following KISS (Keep It Simple Stupid) and code-once deploy-everywhere principles.
* Ethereum interaction learning from scratch. The concept of address, key, signature, etc.
* Add direct interaction from low resources embedded devices that cannot directly sign transactions, such as ESP32, Arduino Uno, etc.
* Add direct interaction from legacy software written in no compatible Ethereum SDK, leveraging the interaction in a secure REST API.
* Connect to the Blockchain some of the most used corporate CRM, ERP, CMS, etc.
* Debug software interaction, and check for errors before publishing your application or service.

## Latest release

!!! Version
    Etherniti Project is under continuous development. Contact development team for more information.
