pragma solidity ^0.7.0;

import "./utils/Administration.sol";
import "./wXMR.sol";

contract Reserve is Administration, wXMR {

    uint256 public              lastProofHeight = 0;
    uint256 public              exchangeRateUSD = 0; // denotes the current USD exchange rate for 1 XMR

    struct Proof {
        bytes signature;
        bytes message;
        bytes reserveAddress;
        uint256 moneroBlockHeight;
        uint256 totalSupply;
    }

    mapping(uint256 => Proof) public proofs;

    event ProofSubmitted(bytes _signature, uint256 _blockNumber);

    constructor() {
        owner = msg.sender;
        admin = msg.sender;
        priceFeed = msg.sender;
    }

    function postReserveProof(
        bytes memory _proof,
        bytes memory _message,
        bytes memory _reserveAddress,
        uint256 _moneroBlockHeight)
        public
        onlyAdmin
        returns (bool)
    {
        proofs[block.number] = Proof({
            signature: _proof,
            message: _message,
            reserveAddress: _reserveAddress,
            moneroBlockHeight: _moneroBlockHeight,
            totalSupply: totalSupply
        });
        lastProofHeight = block.number;
        emit ProofSubmitted(_proof, block.number);
        return true;
    }

    /** @notice used to set the current USD exchange rate for XMR
     */
    function setExchangeRate(
        uint256 _rate
    )
        public
        onlyPriceFeed
        returns (bool)
    {
        exchangeRateUSD = _rate;
        return true;
    }

}