// SPDX-License-Identifier: MIT
pragma solidity >0.4.24;

contract MinerReward {
	struct Miner{
		address		coinbase;
		bool		isMining;
		uint		blockNumber;	
	}


  event CreateMiner(address coinbase, bool isMining, uint blockNumber );

  string public version;
  mapping (address => Miner) public miners;
  
  
  constructor() {
    version = "0.1.0";
  }
  
   function getVersion()public view returns (string memory){
    return version;
  }

  function createMiner(address coinbase) external {
			Miner memory newMiner = Miner(coinbase, true, block.number);
			miners[coinbase]=newMiner;


  }
  
  function stopMining(address coinbase) external{
	Miner storage theMiner = miners[coinbase];
	theMiner.isMining = false;
  }
  
  
}