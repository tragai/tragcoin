# TRAG Token API Reference

This document provides detailed technical reference for interacting with the TRAG token smart contract.

## Contract Information

- **Contract Address**: `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`
- **Network**: Binance Smart Chain (BSC)
- **Standard**: BEP-20 (ERC-20 Compatible)
- **Compiler**: Solidity v0.8.26

## Contract Interface

### View Functions

#### `name() -> string`
Returns the token name.
```solidity
function name() external pure returns (string memory);
```
**Returns**: `"TRAG Coin"`

#### `symbol() -> string`
Returns the token symbol.
```solidity
function symbol() external pure returns (string memory);
```
**Returns**: `"TRAG"`

#### `decimals() -> uint8`
Returns the number of decimals.
```solidity
function decimals() external pure returns (uint8);
```
**Returns**: `6`

#### `totalSupply() -> uint256`
Returns the total token supply.
```solidity
function totalSupply() external view returns (uint256);
```
**Returns**: Total supply in wei (with 6 decimals)

#### `balanceOf(address) -> uint256`
Returns the token balance of an account.
```solidity
function balanceOf(address account) external view returns (uint256);
```
**Parameters**:
- `account`: Address to query balance for

**Returns**: Balance in wei (with 6 decimals)

#### `allowance(address, address) -> uint256`
Returns the remaining tokens that spender can spend on behalf of owner.
```solidity
function allowance(address owner, address spender) external view returns (uint256);
```
**Parameters**:
- `owner`: Token owner address
- `spender`: Address authorized to spend

**Returns**: Remaining allowance in wei

#### `getOwner() -> address`
Returns the contract owner address (BEP-20 specific).
```solidity
function getOwner() external view returns (address);
```

### State-Changing Functions

#### `transfer(address, uint256) -> bool`
Transfers tokens to a specified address.
```solidity
function transfer(address recipient, uint256 amount) external returns (bool);
```
**Parameters**:
- `recipient`: Destination address
- `amount`: Amount in wei (with 6 decimals)

**Requirements**:
- Recipient cannot be zero address
- Caller must have sufficient balance

**Events**: Emits `Transfer` event

#### `approve(address, uint256) -> bool`
Sets allowance for spender with front-running protection.
```solidity
function approve(address spender, uint256 amount) external returns (bool);
```
**Parameters**:
- `spender`: Address authorized to spend
- `amount`: Amount in wei to approve

**Requirements**:
- Spender cannot be zero address
- Cannot approve from non-zero to non-zero (must reset to 0 first)

**Events**: Emits `Approval` event

#### `transferFrom(address, address, uint256) -> bool`
Transfers tokens using allowance mechanism.
```solidity
function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
```
**Parameters**:
- `sender`: Address to transfer from
- `recipient`: Address to transfer to
- `amount`: Amount in wei to transfer

**Requirements**:
- Sender and recipient cannot be zero address
- Sender must have sufficient balance
- Caller must have sufficient allowance

**Events**: Emits `Transfer` and `Approval` events

#### `increaseAllowance(address, uint256) -> bool`
Increases allowance for spender.
```solidity
function increaseAllowance(address spender, uint256 addedValue) external returns (bool);
```
**Parameters**:
- `spender`: Address to increase allowance for
- `addedValue`: Amount to add to current allowance

#### `decreaseAllowance(address, uint256) -> bool`
Decreases allowance for spender.
```solidity
function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool);
```
**Parameters**:
- `spender`: Address to decrease allowance for
- `subtractedValue`: Amount to subtract from current allowance

**Requirements**:
- Current allowance must be >= subtractedValue

### Owner-Only Functions

#### `burn(uint256) -> bool`
Burns tokens from owner's balance (reduces total supply).
```solidity
function burn(uint256 amount) external onlyOwner returns (bool);
```
**Parameters**:
- `amount`: Amount in wei to burn

**Requirements**:
- Only contract owner can call
- Owner must have sufficient balance

## Events

### Transfer
```solidity
event Transfer(address indexed from, address indexed to, uint256 value);
```
Emitted when tokens are transferred.

### Approval
```solidity
event Approval(address indexed owner, address indexed spender, uint256 value);
```
Emitted when allowance is set.

### OwnershipTransferred
```solidity
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
```
Emitted when ownership is transferred.

## Custom Errors

The contract uses custom errors for gas optimization:

- `ZeroAddress()`: Address cannot be zero
- `InsufficientBalance()`: Insufficient balance for operation
- `NonZeroToNonZeroApprove()`: Cannot approve from non-zero to non-zero
- `InsufficientAllowance()`: Insufficient allowance for transfer
- `RenounceDisabled()`: Ownership renouncement is disabled
- `BurnExceedsBalance()`: Burn amount exceeds balance
- `TransferExceedsBalance()`: Transfer amount exceeds balance
- `ApprovalBelowZero()`: Cannot decrease allowance below zero

## Gas Estimates

Approximate gas costs for common operations:

| Function | Gas Estimate |
|----------|--------------|
| transfer | ~51,000 |
| approve | ~46,000 |
| transferFrom | ~56,000 |
| increaseAllowance | ~46,000 |
| decreaseAllowance | ~46,000 |

*Gas costs may vary based on network conditions and wallet states*

## Security Features

1. **Zero Address Protection**: All functions validate against zero addresses
2. **Front-running Protection**: Approve function prevents front-running attacks
3. **Ownership Security**: Renouncement is disabled to maintain control
4. **Custom Errors**: Gas-efficient error handling
5. **Overflow Protection**: Built-in Solidity 0.8.26 overflow protection

## Integration Examples

For complete integration examples in various programming languages, see:
- [JavaScript/Node.js](../examples/javascript/)
- [Python](../examples/python/)
- [Go](../examples/go/)
- [PHP](../examples/php/)