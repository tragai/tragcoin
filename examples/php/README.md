# TRAG Token - PHP Web3 Examples

This directory contains PHP examples for interacting with the TRAG token contract on Binance Smart Chain using Web3.php.

## Prerequisites

- PHP 7.4 or higher
- Composer

## Installation

1. Navigate to this directory:
```bash
cd examples/php
```

2. Install dependencies:
```bash
composer install
```

## Usage

### Basic Setup

```php
<?php
require_once 'vendor/autoload.php';
require_once 'TragWeb3.php';

// For read-only operations (no private key needed)
$trag = new TragToken();

// For transactions (private key required)
$tragWithKey = new TragToken('your_private_key_here');
?>
```

### Get Token Information

```php
$tokenInfo = $trag->getTokenInfo();
print_r($tokenInfo);

// Output:
// Array
// (
//     [name] => TRAG Coin
//     [symbol] => TRAG
//     [decimals] => 6
//     [totalSupply] => 100000000000.000000
//     [contractAddress] => 0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C
// )
```

### Check Token Balance

```php
$balance = $trag->getBalance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8');
echo "Balance: " . $balance['formatted'] . " TRAG\n";
```

### Check Allowance

```php
$allowance = $trag->getAllowance('owner_address', 'spender_address');
echo "Allowance: " . $allowance['formatted'] . " TRAG\n";
```

### Transfer Tokens (Note: Requires additional implementation)

```php
// This requires additional Web3 account implementation
try {
    $result = $tragWithKey->transfer('recipient_address', '100');
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}
```

## Running Examples

Execute the example script:

```bash
php TragWeb3.php
```

Or use composer:

```bash
composer run example
```

## Error Handling

The TragToken class includes comprehensive error handling:

```php
try {
    $balance = $trag->getBalance('invalid_address');
} catch (Exception $e) {
    echo "Error: " . $e->getMessage() . "\n";
}
```

## Number Precision

This library uses BCMath for precise decimal calculations:

```php
// Converting amounts
$amountWei = bcmul('100.123456', bcpow('10', '6')); // To wei
$amountFormatted = bcdiv($amountWei, bcpow('10', '6'), 6); // From wei
```

## Environment Variables

For production use, store sensitive data in environment variables:

```php
$privateKey = $_ENV['PRIVATE_KEY'] ?? '';
$trag = new TragToken($privateKey);
```

## Security Notes

- Never commit private keys to version control
- Use environment variables for sensitive data
- Always validate addresses before sending transactions
- Test transactions with small amounts first
- Enable BCMath extension for precise calculations

## Contract Address

**TRAG Token Contract:** `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`

## TRAG Ecosystem Addresses

### Official Wallets
- **TRAG Investor:** `0x7E18Dc9e51965aBC7Db6F9c41Abeb040122400C8`
- **TRAG Team:** `0xB57B316BC2eE445c90e2F4e04DA6D6da1b53cf05`
- **TRAG Rewards:** `0x9a5f69e1bb111789059CaeE4A156FC490e46DC4e`
- **TRAG Treasury:** `0x30B252B746C49651f822d07b12B31B2d73e56af8`
- **TRAG Marketing:** `0xd2839de718cbBe5E13fD4DEFae92010fE6F26383`
- **TRAG Liquidity:** `0x1Cab22De64a3c852AC81fAd9a1590600eCd02325`

## BSC Network Configuration

- **Network:** Binance Smart Chain
- **Chain ID:** 56
- **RPC URL:** https://bsc-dataseed1.binance.org/
- **Explorer:** https://bscscan.com/

## Dependencies

- **Web3.php**: PHP library for interacting with Ethereum-compatible blockchains
- **BCMath**: PHP extension for arbitrary precision mathematics

## Limitations

This example provides a basic framework for interacting with TRAG token. Full transaction functionality requires additional implementation for:

- Private key management
- Transaction signing
- Gas estimation
- Nonce management

For production applications, consider using more robust Web3 PHP libraries or services.