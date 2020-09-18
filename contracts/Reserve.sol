pragma solidity ^0.7.0;

import "./utils/Administration.sol";
import "./wXMR.sol";

contract Reserve is Administration, wXMR {

    uint256 public  lastProofHeight = 0;
    uint256 public  exchangeRateUSD = 0; // denotes the current USD exchange rate for 1 XMR
    bytes   public  reserveAddress; // address of the monero wallet functioning as a reserve

    struct Proof {
        bytes32 signatureKeccakHash;
        bytes message;
        bytes reserveAddress;
        uint256 moneroBlockHeight;
        uint256 totalSupply;
    }

    mapping(uint256 => Proof) public proofs;

    event ProofSubmitted(bytes32 _signatureKeccakHash, uint256 _blockNumber);

    constructor(bytes memory _reserveAddress) {
        owner = msg.sender;
        admin = msg.sender;
        priceFeed = msg.sender;
        reserveAddress = _reserveAddress;
    }

    function postReserveProof(
        bytes memory _message,
        bytes memory _reserveAddress,
        bytes32 _signatureKeccakHash,
        uint256 _moneroBlockHeight)
        public
        onlyAdmin
        returns (bool)
    {
        proofs[block.number] = Proof({
            signatureKeccakHash: _signatureKeccakHash,
            message: _message,
            reserveAddress: _reserveAddress,
            moneroBlockHeight: _moneroBlockHeight,
            totalSupply: totalSupply
        });
        lastProofHeight = block.number;
        emit ProofSubmitted(_signatureKeccakHash, block.number);
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