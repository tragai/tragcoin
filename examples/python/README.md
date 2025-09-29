# TRAG Token - Python Web3 Examples

This directory contains Python examples for interacting with the TRAG token contract on Binance Smart Chain using Web3.py.

## Prerequisites

- Python 3.7 or higher
- pip

## Installation

1. Navigate to this directory:
```bash
cd examples/python
```

2. Install dependencies:
```bash
pip install -r requirements.txt
```

## Usage

### Basic Setup

```python
from trag_web3 import TragToken

# For read-only operations (no private key needed)
trag = TragToken()

# For transactions (private key required)
trag_with_key = TragToken('your_private_key_here')
```

### Get Token Information

```python
token_info = trag.get_token_info()
print(token_info)
# Output:
# {
#   'name': 'TRAG Coin',
#   'symbol': 'TRAG',
#   'decimals': 6,
#   'total_supply': 100000000000.0,
#   'contract_address': '0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C'
# }
```

### Check Token Balance

```python
balance = trag.get_balance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8')
print(f"Balance: {balance['formatted']} TRAG")
```

### Transfer Tokens

```python
# Requires private key
result = trag_with_key.transfer('recipient_address', 100)
print(f"Transfer successful: {result['transaction_hash']}")
```

### Approve Tokens

```python
# Requires private key
result = trag_with_key.approve('spender_address', 1000)
print(f"Approval successful: {result['transaction_hash']}")
```

### Check Allowance

```python
allowance = trag.get_allowance('owner_address', 'spender_address')
print(f"Allowance: {allowance['formatted']} TRAG")
```

## Running Examples

Execute the example script:

```bash
python trag_web3.py
```

## Error Handling

The TragToken class includes comprehensive error handling:

```python
try:
    balance = trag.get_balance('invalid_address')
except Exception as e:
    print(f"Error: {str(e)}")
```

## Environment Variables

For production use, store sensitive data in environment variables:

```python
import os
from trag_web3 import TragToken

private_key = os.getenv('PRIVATE_KEY')
trag = TragToken(private_key)
```

## Security Notes

- Never commit private keys to version control
- Use environment variables for sensitive data
- Always validate addresses before sending transactions
- Test transactions with small amounts first

## Contract Address

**TRAG Token Contract:** `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`

## BSC Network Configuration

- **Network:** Binance Smart Chain
- **Chain ID:** 56
- **RPC URL:** https://bsc-dataseed1.binance.org/
- **Explorer:** https://bscscan.com/

## Dependencies

- **web3**: Python library for interacting with Ethereum-compatible blockchains