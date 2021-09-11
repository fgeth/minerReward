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
  mapping (address => Miner) public Miners;
  address[] public Addresses;
  
  
  constructor() {
    version = "0.1.0";
  }
  
   function getVersion()public view returns (string memory){
    return version;
  }

  function createMiner(address coinbase) external {
			Miner memory newMiner = Miner(coinbase, true, block.number);
			Miners[coinbase]=newMiner;
			bool minerExists = false;
			for (uint i = 0; i < Addresses.length; i++) {
				if (Addresses[i]==coinbase){
					minerExists =true;
				}
			}
			if(!minerExists){
				Addresses.push(coinbase);
			}


  }
  
  function stopMining(address coinbase) external{
	Miners[coinbase].isMining = false;
  }
  
  function getMiners() public view returns (address[] memory){
		return Addresses;
  }
  
}