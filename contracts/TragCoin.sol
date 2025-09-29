/**
 * Submitted for verification at BscScan.com
 */

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

// BEP20 Token Standard Interface
interface IBEP20 {
  /**
   * @dev Returns the amount of tokens in existence
   */
  function totalSupply() external view returns (uint256);

  /**
   * @dev Returns the token decimals
   */
  function decimals() external view returns (uint8);

  /**
   * @dev Returns the token symbol
   */
  function symbol() external view returns (string memory);

  /**
   * @dev Returns the token name
   */
  function name() external view returns (string memory);

  /**
   * @dev Returns the bep token owner
   */
  function getOwner() external view returns (address);

  /**
   * @dev Returns the amount of tokens owned by `account`
   * @param account The address to query the balance of
   */
  function balanceOf(address account) external view returns (uint256);

  /**
   * @dev Moves `amount` tokens from the caller's account to `recipient`
   * @param recipient The address to transfer to
   * @param amount The amount to be transferred
   * @return Returns a boolean value indicating whether the operation succeeded
   *
   * Emits a Transfer event
   */
  function transfer(address recipient, uint256 amount) external returns (bool);

  /**
   * @dev Returns the remaining number of tokens that `spender` will be allowed to spend on behalf of `owner`
   * @param _owner The address which owns the funds
   * @param spender The address which will spend the funds
   * @return The remaining allowance
   */
  function allowance(address _owner, address spender) external view returns (uint256);

  /**
   * @dev Sets `amount` as the allowance of `spender` over the caller's tokens
   * @param spender The address which will spend the funds
   * @param amount The amount of tokens to be spent
   * @return Returns a boolean value indicating whether the operation succeeded
   *
   * Emits an Approval event
   */
  function approve(address spender, uint256 amount) external returns (bool);

  /**
   * @dev Moves `amount` tokens from `sender` to `recipient` using the allowance mechanism
   * @param sender The address to transfer from
   * @param recipient The address to transfer to
   * @param amount The amount to be transferred
   * @return Returns a boolean value indicating whether the operation succeeded
   *
   * Emits a Transfer event
   */
  function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);

  /**
   * @dev Emitted when `value` tokens are moved from one account (`from`) to another (`to`)
   * @param from The address tokens are transferred from
   * @param to The address tokens are transferred to
   * @param value The amount of tokens transferred
   */
  event Transfer(address indexed from, address indexed to, uint256 value);

  /**
   * @dev Emitted when the allowance of a `spender` for an `owner` is set by a call to approve
   * @param owner The address which owns the funds
   * @param spender The address which will spend the funds
   * @param value The amount of tokens approved
   */
  event Approval(address indexed owner, address indexed spender, uint256 value);
}

// Provides information about the current execution context
contract Context {
  // Empty internal constructor to prevent direct deployment
  constructor () { }

  // Returns the sender of the transaction
  function _msgSender() internal view returns (address) {
    return msg.sender;
  }

  // Returns the data of the transaction
  function _msgData() internal view returns (bytes memory) {
    this; // Silence state mutability warning
    return msg.data;
  }
}

/**
 * @dev Contract module which provides a basic access control mechanism
 * where there is an account (an owner) that can be granted exclusive access to specific functions
 * By default, the owner account will be the one that deploys the contract
 */
contract Ownable is Context {
  address private _owner;

  // Emitted when ownership is transferred
  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

  /**
   * @dev Initializes the contract setting the deployer as the initial owner
   */
  constructor ()  {
    address msgSender = _msgSender();
    _owner = msgSender;
    emit OwnershipTransferred(address(0), msgSender);
  }

  /**
   * @dev Returns the address of the current owner
   */
  function owner() public view returns (address) {
    return _owner;
  }

  /**
   * @dev Throws if called by any account other than the owner
   */
  modifier onlyOwner() {
    require(_owner == _msgSender(), "Ownable: caller is not the owner");
    _;
  }

  /**
   * @dev Leaves the contract without owner. It will not be possible to call onlyOwner functions anymore
   * Can only be called by the current owner
   * NOTE: Renouncing ownership will leave the contract without an owner
   */
  function renounceOwnership() public virtual onlyOwner {
    emit OwnershipTransferred(_owner, address(0));
    _owner = address(0);
  }

  /**
   * @dev Transfers ownership of the contract to a new account
   * @param newOwner The address of the new owner
   */
  function transferOwnership(address newOwner) public onlyOwner {
    _transferOwnership(newOwner);
  }

  /**
   * @dev Transfers ownership of the contract to a new account
   * @param newOwner The address of the new owner
   */
  function _transferOwnership(address newOwner) internal {
    require(newOwner != address(0), "Ownable: new owner is the zero address");
    emit OwnershipTransferred(_owner, newOwner);
    _owner = newOwner;
  }
}

// Custom Errors for gas optimization
error ZeroAddress();
error InsufficientBalance();
error NonZeroToNonZeroApprove();
error InsufficientAllowance();
error RenounceDisabled();
error BurnExceedsBalance();
error TransferExceedsBalance();
error ApprovalBelowZero();

// TRAG Coin BEP20 Token Implementation
contract TRAGCoin is Context, IBEP20, Ownable {

  // Stores balances for each address
  mapping (address => uint256) private _balances;

  // Stores allowances: owner => spender => amount
  mapping (address => mapping (address => uint256)) private _allowances;

  // Total supply of tokens
  uint256 private _totalSupply;
  // Token decimals (6)
  uint8 private constant _decimals = 6;
  // Token symbol
  string private constant _symbol = "TRAG";
  // Token name
  string private constant _name = "TRAG Coin";

  /**
   * @dev Constructor that gives msg.sender all of existing tokens
   * Total supply: 100,000,000,000 * 10^6
   * Emits Transfer event from address(0) to msg.sender
   */
  constructor() {
    _totalSupply = 100_000_000_000 * 10 ** uint256(_decimals); // 100 billion tokens
    _balances[msg.sender] = _totalSupply;

    emit Transfer(address(0), msg.sender, _totalSupply);
  }

  /**
   * @dev Returns the bep token owner
   */
  function getOwner() external view returns (address) {
    return owner();
  }

  /**
   * @dev Returns the token decimals
   */
  function decimals() external pure returns (uint8) {
    return _decimals;
  }

  /**
   * @dev Returns the token symbol
   */
  function symbol() external pure returns (string memory) {
    return _symbol;
  }

  /**
   * @dev Returns the token name
   */
  function name() external pure returns (string memory) {
    return _name;
  }

  /**
   * @dev Returns the total token supply
   */
  function totalSupply() external view returns (uint256) {
    return _totalSupply;
  }

  /**
   * @dev Returns balance of the given account
   * @param account The address to query
   */
  function balanceOf(address account) external view returns (uint256) {
    return _balances[account];
  }

  /**
   * @dev Transfer tokens to a specified address
   * @param recipient The address to transfer to
   * @param amount The amount to be transferred
   * @return Success boolean
   *
   * Requirements:
   * - recipient cannot be the zero address
   * - the caller must have sufficient balance
   */
  function transfer(address recipient, uint256 amount) external returns (bool) {
    _transfer(_msgSender(), recipient, amount);
    return true;
  }

  /**
   * @dev Returns allowance for spender to spend owner's tokens
   * @param owner The token owner address
   * @param spender The spender address
   */
  function allowance(address owner, address spender) external view returns (uint256) {
    return _allowances[owner][spender];
  }

  /**
   * @dev Approve spender to spend amount of caller's tokens
   * @param spender The address authorized to spend
   * @param amount The max amount they can spend
   * @return Success boolean
   *
   * Requirements:
   * - spender cannot be the zero address
   * - prevents approve from non-zero to non-zero (front-running protection)
   */
  function approve(address spender, uint256 amount) external returns (bool) {
    uint256 cur = _allowances[_msgSender()][spender];
    if (cur != 0 && amount != 0) revert NonZeroToNonZeroApprove();
    _approve(_msgSender(), spender, amount);
    return true;
  }

  // safeApprove functionality is now integrated into approve()
  // The approve function performs the same security validation

  /**
   * @dev Transfer from one address to another using allowance
   * @param sender The address to send from
   * @param recipient The address to receive
   * @param amount The amount to transfer
   * @return Success boolean
   *
   * Requirements:
   * - sender and recipient cannot be zero address
   * - sender must have sufficient balance
   * - caller must have allowance for sender's tokens >= amount
   *
   * Emits Transfer and Approval events
   */
  function transferFrom(address sender, address recipient, uint256 amount) external returns (bool) {
    uint256 currentAllowance = _allowances[sender][_msgSender()];
    if (currentAllowance < amount) revert InsufficientAllowance();

    _transfer(sender, recipient, amount);
    _approve(sender, _msgSender(), currentAllowance - amount);
    return true;
  }

  /**
   * @dev Increase the allowance granted to spender
   * @param spender The address authorized to spend
   * @param addedValue The amount to increase allowance by
   * @return Success boolean
   *
   * Emits Approval event
   */
  function increaseAllowance(address spender, uint256 addedValue) public returns (bool) {
    _approve(_msgSender(), spender, _allowances[_msgSender()][spender] + addedValue);
    return true;
  }

  /**
   * @dev Decrease the allowance granted to spender
   * @param spender The address authorized to spend
   * @param subtractedValue The amount to decrease allowance by
   * @return Success boolean
   *
   * Requirements:
   * - current allowance must be >= subtractedValue
   * Emits Approval event
   */
  function decreaseAllowance(address spender, uint256 subtractedValue) public returns (bool) {
    uint256 currentAllowance = _allowances[_msgSender()][spender];
    if (currentAllowance < subtractedValue) revert ApprovalBelowZero();

    _approve(_msgSender(), spender, currentAllowance - subtractedValue);
    return true;
  }

  /**
   * @dev Internal transfer function
   * @param sender The address sending tokens
   * @param recipient The address receiving tokens
   * @param amount The amount to transfer
   *
   * Requirements:
   * - sender and recipient cannot be zero address
   * - sender must have sufficient balance
   * Emits Transfer event
   */
  function _transfer(address sender, address recipient, uint256 amount) internal {
    if (sender == address(0)) revert ZeroAddress();
    if (recipient == address(0)) revert ZeroAddress();
    if (_balances[sender] < amount) revert TransferExceedsBalance();

    _balances[sender] -= amount;
    _balances[recipient] += amount;
    emit Transfer(sender, recipient, amount);
  }

  /**
   * @dev Internal burn function
   * @param account The address to burn from
   * @param amount The amount to burn
   *
   * Requirements:
   * - account cannot be zero address
   * - account must have sufficient balance
   * Emits Transfer event to address(0)
   */
  function _burn(address account, uint256 amount) internal {
    if (account == address(0)) revert ZeroAddress();
    if (_balances[account] < amount) revert BurnExceedsBalance();

    _balances[account] -= amount;
    _totalSupply -= amount;
    emit Transfer(account, address(0), amount);
  }

  /**
   * @dev Internal approve function
   * @param owner Token owner address
   * @param spender Spender address
   * @param amount Amount to approve
   *
   * Requirements:
   * - owner and spender cannot be zero address
   * Emits Approval event
   */
  function _approve(address owner, address spender, uint256 amount) internal {
    if (owner == address(0)) revert ZeroAddress();
    if (spender == address(0)) revert ZeroAddress();

    _allowances[owner][spender] = amount;
    emit Approval(owner, spender, amount);
  }

  /**
   * @dev Ownership renouncement is disabled for security
   * Owner must always exist to maintain token management functions
   * Always reverts when called
   */
  function renounceOwnership() public view override onlyOwner {
    revert RenounceDisabled();
  }

  /**
   * @dev Burns tokens from owner's balance
   * @param amount Amount to burn
   * @return Success boolean
   */
  function burn(uint256 amount) public onlyOwner returns (bool) {
    _burn(_msgSender(), amount);
    return true;
  }
}