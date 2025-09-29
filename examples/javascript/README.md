# TRAG Token - JavaScript/Node.js Web3 Examples

This directory contains JavaScript examples for interacting with the TRAG token contract on Binance Smart Chain.

## Prerequisites

- Node.js (version 16 or higher)
- npm or yarn

## Installation

1. Navigate to this directory:
```bash
cd examples/javascript
```

2. Install dependencies:
```bash
npm install
```

## Usage

### Basic Setup

```javascript
const { TragToken } = require('./trag-web3');

// For read-only operations (no private key needed)
const trag = new TragToken();

// For transactions (private key required)
const tragWithKey = new TragToken('your_private_key_here');
```

### Get Token Information

```javascript
const tokenInfo = await trag.getTokenInfo();
console.log(tokenInfo);
// Output:
// {
//   name: 'TRAG Coin',
//   symbol: 'TRAG',
//   decimals: 6,
//   totalSupply: '100000000000',
//   contractAddress: '0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C'
// }
```

### Check Token Balance

```javascript
const balance = await trag.getBalance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8');
console.log(`Balance: ${balance.formatted} TRAG`);
```

### Transfer Tokens

```javascript
// Requires private key
const result = await tragWithKey.transfer('recipient_address', '100');
console.log('Transfer successful:', result.transactionHash);
```

### Approve Tokens

```javascript
// Requires private key
const result = await tragWithKey.approve('spender_address', '1000');
console.log('Approval successful:', result.transactionHash);
```

### Check Allowance

```javascript
const allowance = await trag.getAllowance('owner_address', 'spender_address');
console.log(`Allowance: ${allowance.formatted} TRAG`);
```

## Running Examples

Execute the example script:

```bash
npm start
# or
node trag-web3.js
```

## Security Notes

- Never commit private keys to version control
- Use environment variables for sensitive data
- Always validate addresses before sending transactions
- Test on BSC testnet before mainnet transactions

## Contract Address

**TRAG Token Contract:** `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`

## BSC Network Configuration

- **Network:** Binance Smart Chain
- **Chain ID:** 56
- **RPC URL:** https://bsc-dataseed1.binance.org/
- **Explorer:** https://bscscan.com/