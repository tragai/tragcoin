# TRAG Token API Reference

This document provides detailed technical reference for interacting with the TRAG token smart contract.

## Contract Information

- **Contract Address**: `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`
- **Network**: Binance Smart Chain (BSC)
- **Standard**: BEP-20 (ERC-20 Compatible)
- **Compiler**: Solidity v0.8.26+commit.8a97fa7a
- **Deployed Bytecode**: `0x608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063a457c2d711610066578063a457c2d714610242578063a9059cbb14610255578063dd62ed3e14610268578063f2fde38b146102a157600080fd5b8063715018a6146101e2578063893d20e8146101ec5780638da5cb5b1461021157806395d89b411461022257600080fd5b8063313ce567116100d3578063313ce56714610184578063395093511461019357806342966c68146101a657806370a08231146101b957600080fd5b806306fdde0314610105578063095ea7b31461013c57806318160ddd1461015f57806323b872dd14610171575b600080fd5b6040805180820190915260098152682a2920a39021b7b4b760b91b60208201525b604051610133919061086e565b60405180910390f35b61014f61014a3660046108d3565b6102b4565b6040519015158152602001610133565b6003545b604051908152602001610133565b61014f61017f3660046108fd565b61031b565b60405160068152602001610133565b61014f6101a13660046108d3565b61038a565b61014f6101b436600461093a565b6103ca565b6101636101c7366004610953565b6001600160a01b031660009081526001602052604090205490565b6101ea610411565b005b6000546001600160a01b03165b6040516001600160a01b039091168152602001610133565b6000546001600160a01b03166101f9565b6040805180820190915260048152635452414760e01b6020820152610126565b61014f6102503660046108d3565b610454565b61014f6102633660046108d3565b6104a8565b610163610276366004610975565b6001600160a01b03918216600090815260026020908152604080832093909416825291909152205490565b6101ea6102af366004610953565b6104b5565b3360009081526002602090815260408083206001600160a01b038616845290915281205480158015906102e657508215155b1561030457604051630a65468f60e11b815260040160405180910390fd5b61030f3385856104eb565b60019150505b92915050565b6001600160a01b038316600090815260026020908152604080832033845290915281205482811015610360576040516313be252b60e01b815260040160405180910390fd5b61036b85858561059b565b61037f853361037a86856109be565b6104eb565b506001949350505050565b3360008181526002602090815260408083206001600160a01b038716845290915281205490916103c191859061037a9086906109d1565b50600192915050565b600080546001600160a01b031633146103fe5760405162461bcd60e51b81526004016103f5906109e4565b60405180910390fd5b61040833836106c3565b5060015b919050565b6000546001600160a01b0316331461043b5760405162461bcd60e51b81526004016103f5906109e4565b604051638905116560e01b815260040160405180910390fd5b3360009081526002602090815260408083206001600160a01b0386168452909152812054828110156104995760405163dca94c6160e01b815260040160405180910390fd5b61030f338561037a86856109be565b60006103c133848461059b565b6000546001600160a01b031633146104df5760405162461bcd60e51b81526004016103f5906109e4565b6104e8816107ae565b50565b6001600160a01b0383166105125760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0382166105395760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0383811660008181526002602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b0383166105c25760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0382166105e95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0383166000908152600160205260409020548111156106225760405163169b037b60e01b815260040160405180910390fd5b6001600160a01b0383166000908152600160205260408120805483929061064a9084906109be565b90915550506001600160a01b038216600090815260016020526040812080548392906106779084906109d1565b92505081905550816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161058e91815260200190565b6001600160a01b0382166106ea5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0382166000908152600160205260409020548111156107235760405163588569f760e01b815260040160405180910390fd5b6001600160a01b0382166000908152600160205260408120805483929061074b9084906109be565b92505081905550806003600082825461076491906109be565b90915550506040518181526000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6001600160a01b0381166108135760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103f5565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b602081526000825180602084015260005b8181101561089c576020818601810151604086840101520161087f565b506000604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b038116811461040c57600080fd5b600080604083850312156108e657600080fd5b6108ef836108bc565b946020939093013593505050565b60008060006060848603121561091257600080fd5b61091b846108bc565b9250610929602085016108bc565b929592945050506040919091013590565b60006020828403121561094c57600080fd5b5035919050565b60006020828403121561096557600080fd5b61096e826108bc565b9392505050565b6000806040838503121561098857600080fd5b610991836108bc565b915061099f602084016108bc565b90509250929050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610315576103156109a8565b80820180821115610315576103156109a8565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260408201526060019056fea2646970667358221220b546875c219e7c52a8b388235627493ae015c16b8f49d99137b32fb656a99b0364736f6c634300081a0033`

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