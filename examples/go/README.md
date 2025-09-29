# TRAG Token - Go Web3 Examples

This directory contains Go examples for interacting with the TRAG token contract on Binance Smart Chain using go-ethereum.

## Prerequisites

- Go 1.19 or higher
- Git

## Installation

1. Navigate to this directory:
```bash
cd examples/go
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

### Basic Setup

```go
import "path/to/your/project/trag_web3"

// For read-only operations (no private key needed)
trag, err := NewTragToken("")
if err != nil {
    log.Fatal(err)
}

// For transactions (private key required)
tragWithKey, err := NewTragToken("your_private_key_hex")
if err != nil {
    log.Fatal(err)
}
```

### Get Token Information

```go
tokenInfo, err := trag.GetTokenInfo()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Name: %s\n", tokenInfo.Name)
fmt.Printf("Symbol: %s\n", tokenInfo.Symbol)
fmt.Printf("Decimals: %d\n", tokenInfo.Decimals)
fmt.Printf("Total Supply: %s\n", tokenInfo.TotalSupply.String())
```

### Check Token Balance

```go
balance, err := trag.GetBalance("0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Balance: %s TRAG\n", balance.Formatted.String())
```

### Transfer Tokens

```go
// Requires private key
amount := big.NewFloat(100) // 100 TRAG
tx, err := tragWithKey.Transfer("recipient_address", amount)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Transfer successful: %s\n", tx.Hash().Hex())
```

### Check Allowance

```go
allowance, err := trag.GetAllowance("owner_address", "spender_address")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Allowance: %s TRAG\n", allowance.Formatted.String())
```

## Running Examples

Execute the example program:

```bash
go run trag_web3.go
```

Or build and run:

```bash
go build -o trag-example trag_web3.go
./trag-example
```

## Error Handling

All functions return errors that should be handled appropriately:

```go
balance, err := trag.GetBalance("invalid_address")
if err != nil {
    log.Printf("Error getting balance: %v", err)
    return
}
```

## Environment Variables

For production use, store sensitive data in environment variables:

```go
import "os"

privateKey := os.Getenv("PRIVATE_KEY")
trag, err := NewTragToken(privateKey)
```

## Big Number Handling

This library uses Go's `big.Int` and `big.Float` for precise decimal calculations:

```go
// Converting from string
amount, _ := big.NewFloat(0).SetString("123.456789")

// Converting to string
amountStr := amount.String()

// Mathematical operations
result := new(big.Float).Add(amount1, amount2)
```

## Security Notes

- Never commit private keys to version control
- Use environment variables for sensitive data
- Always validate addresses before sending transactions
- Test transactions with small amounts first
- Private keys should be provided without the "0x" prefix

## Contract Address

**TRAG Token Contract:** `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`

## BSC Network Configuration

- **Network:** Binance Smart Chain
- **Chain ID:** 56
- **RPC URL:** https://bsc-dataseed1.binance.org/
- **Explorer:** https://bscscan.com/

## Dependencies

- **go-ethereum**: Official Go implementation of the Ethereum protocol
- **big**: Go's arbitrary precision arithmetic package