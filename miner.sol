// SPDX-License-Identifier: MIT
pragma solidity >0.4.24;

contract MinerReward {
	struct Miner{
		address		coinbase;
		bool		isMining ==false;
		uint		blockNumber;	
	}


  event CreateMiner(address coinbase, bool isMining, uint blockNumber );

  string public version;
  Miner[] public miners;

  constructor() {
    version = "0.1.0";
  }
  
   function getVersion()public pure returns (string memory){
    return version;
  }

  function createMiner(address memory coinbase) external {
			var Miner newMiner = coinbase;
			newMiner.isMining = true;
			newMiner.blockNumber = block.number;
			miners[coinbase]=newMiner;


  }
  
  function stopMining(address memory coinbase){
	var theMiner = miners[coinbase];
	theMiner.isMining = false;
  }
  
  function getMiners() public view returns (Miner[] memory){
    return miners;
  }
  
}