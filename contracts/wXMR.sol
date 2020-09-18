pragma solidity ^0.7.0;

import "./utils/SafeMath.sol";
import "./utils/Administration.sol";

contract wXMR is Administration {

    using SafeMath for uint256;

    string  constant public  name = "wXMR";
    string  constant public  symbol = "wXMR";
    uint256 public  totalSupply = 0;
    uint256 public exchangeRateUSD = 0; // denotes the current USD exchange rate for 1 XMR
    uint8   constant public  decimals = 12;
    bytes   public latestReserveProof;

    mapping (address => uint256) private balances;
    mapping (address => mapping (address => uint256)) private allowed;

    event Transfer(address indexed _sender, address indexed _recipient, uint256 _amount);
    event Approval(address indexed _owner, address indexed _spender, uint256 _amount);
    event CoinsMinted(address indexed _recipient, uint256 _mintAmount);
    event CoinsBurned(address indexed _burner, uint256 _burnAmount, bytes32 _moneroAddressHash);
    event ReserveProofUpdate(bytes _new, bytes _previous);

    constructor() {
        owner = msg.sender;
        admin = msg.sender;
    }

    /** @notice updates a reserve proof from the multisig wallet
    */
    function postReserveProof(
        bytes memory _proof)
        public
        onlyAdmin
        returns (bool)
    {
        bytes memory oldProof = latestReserveProof;
        latestReserveProof = _proof;
        emit ReserveProofUpdate(_proof, oldProof);
        return true;
    }

    /** @notice used to set the current USD exchange rate for XMR
     */
    function setExchangeRate(
        uint256 _rate
    )
        public
        returns (bool)
    {
        exchangeRateUSD = _rate;
        return true;
    }

    /** @notice Used to transfer tokens
        * @param _recipient This is the recipient of the transfer
        * @param _amount This is the amount of tokens to send
     */
    function transfer(
        address _recipient,
        uint256 _amount
    )
        public
        nonZeroAddress(_recipient)
        returns (bool)
    {
        // dont check require as this will throw getting the same effect 
        balances[msg.sender] = balances[msg.sender].sub(_amount);
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(msg.sender, _recipient, _amount);
        return true;
    }

    /** @notice Used to transfer tokens on behalf of someone else
        * @param _recipient This is the recipient of the transfer
        * @param _amount This is the amount of tokens to send
     */
    function transferFrom(
        address _owner,
        address _recipient,
        uint256 _amount
    )
        public
        nonZeroAddress(_recipient)
        returns (bool)
    {
        // reduce the allowance
        allowed[_owner][msg.sender] = allowed[_owner][msg.sender].sub(_amount);
        // reduce balance of owner
        balances[_owner] = balances[_owner].sub(_amount);
        // increase balance of recipient
        balances[_recipient] = balances[_recipient].add(_amount);
        emit Transfer(_owner, _recipient, _amount);
        return true;
    }

    /** @notice This is used to approve someone to send tokens on your behalf
        * @param _spender This is the person who can spend on your behalf
        * @param _value This is the amount of tokens that they can spend
     */
    function approve(address _spender, uint256 _value) public returns (bool) {
        allowed[msg.sender][_spender] = _value;
        emit Approval(msg.sender, _spender, _value);
        return true;
    }

    /** @notice This is used to mint new tokens
        * @dev Can only be executed by the staking, and merged miner validator contracts
        * @param _recipient This is the person who will received the mint tokens
        * @param _amount This is the amount of tokens that they will receive and which will be generated
     */
    function mint(
        address _recipient,
        uint256 _amount)
        public // todo: restrict access
        onlyAdmin
        returns (bool)
    {
        balances[_recipient] = balances[_recipient].add(_amount);
        totalSupply = totalSupply.add(_amount);
        emit Transfer(address(0), _recipient, _amount);
        emit CoinsMinted(_recipient, _amount);
        return true;
    }

    /** @notice burns wXMR so that the underlying XMR may be redeemed
      * @param _moneroAddressHash the keccak256 hash of the monero address that will receive the funds
      * @param _moneroAddressHash you must present the address to the web portal which will validate the hash, along with the transaction it was burned in
     */
    function burn(
        uint256 _amount,
        bytes32 _moneroAddressHash)
        public
        returns (bool)
    {
        balances[msg.sender] = balances[msg.sender].sub(_amount);
        totalSupply = totalSupply.sub(_amount);
        emit Transfer(msg.sender, address(0), _amount);
        emit CoinsBurned(msg.sender, _amount, _moneroAddressHash);
    }

    /**
    * @dev Increase the amount of tokens that an owner allowed to a spender.
    * approve should be called when allowed[_spender] == 0. To increment
    * allowed value is better to use this function to avoid 2 calls (and wait until
    * the first transaction is mined)
    * From MonolithDAO Token.sol
    * @param _spender The address which will spend the funds.
    * @param _addedValue The amount of tokens to increase the allowance by.
    */
    function increaseApproval(
        address _spender,
        uint256 _addedValue
    )
        public
        returns (bool)
    {
        allowed[msg.sender][_spender] = allowed[msg.sender][_spender].add(_addedValue);
        emit Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
        return true;
    }

    /**
    * @dev Decrease the amount of tokens that an owner allowed to a spender.
    * approve should be called when allowed[_spender] == 0. To decrement
    * allowed value is better to use this function to avoid 2 calls (and wait until
    * the first transaction is mined)
    * From MonolithDAO Token.sol
    * @param _spender The address which will spend the funds.
    * @param _subtractedValue The amount of tokens to decrease the allowance by.
    */
    function decreaseApproval(
        address _spender,
        uint256 _subtractedValue
    )
        public
        returns (bool)
    {
        uint256 oldValue = allowed[msg.sender][_spender];
        if (_subtractedValue >= oldValue) {
            allowed[msg.sender][_spender] = 0;
        } else {
            allowed[msg.sender][_spender] = oldValue.sub(_subtractedValue);
        }
        emit Approval(msg.sender, _spender, allowed[msg.sender][_spender]);
        return true;
    }

    /**GETTERS */

    /** @notice Used to get the balance of a holder
        * @param _holder The address of the token holder
     */
    function balanceOf(
        address _holder
    )
        public
        view
        returns (uint256)
    {
        return balances[_holder];
    }

    /** @notice Used to get the allowance of someone
        * @param _owner The address of the token owner
        * @param _spender The address of thhe person allowed to spend funds on behalf of the owner
     */
    function allowance(
        address _owner,
        address _spender
    )
        public
        view
        returns (uint256)
    {
        return allowed[_owner][_spender];
    }

}
