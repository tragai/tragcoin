# TRAG Coin

![TRAG Coin Logo](logo.png)

TRAG Coin is a BEP-20 token deployed on the Binance Smart Chain (BSC) network.

> **For complete project information, whitepaper, and official updates, visit [tragcoin.com](https://tragcoin.com)**

This repository contains the technical resources for developers integrating with TRAG token.

## Token Information

- **Token Name:** TRAG Coin
- **Symbol:** TRAG
- **Network:** Binance Smart Chain (BSC)
- **Contract Address:** `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`
- **Decimals:** 6
- **Total Supply:** 100,000,000,000 TRAG
- **Standard:** BEP-20 (ERC-20 Compatible)

## Quick Links

- [Official Website & Whitepaper](https://tragcoin.com)
- [Telegram Community](https://t.me/tragcoin)
- [View on BSCScan](https://bscscan.com/address/0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C)
- [Add to MetaMask](#adding-to-metamask)

## Adding to MetaMask

To add TRAG Coin to your MetaMask wallet:

1. Open MetaMask
2. Click "Import tokens"
3. Select "Custom Token"
4. Enter the contract address: `0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C`
5. Token symbol and decimals should auto-fill
6. Click "Add Custom Token"

## BSC Network Configuration

If you haven't added BSC to MetaMask yet:

- **Network Name:** Smart Chain
- **New RPC URL:** https://bsc-dataseed1.binance.org/
- **ChainID:** 56
- **Symbol:** BNB
- **Block Explorer URL:** https://bscscan.com/

## Contract Overview

TRAG Coin is implemented as a comprehensive BEP-20 token with enhanced security features and gas optimizations. The contract follows industry best practices while providing additional functionality for token management.

### Key Features

- **Standard Compliance:** Full BEP-20 (ERC-20 compatible) implementation
- **Burn Functionality:** Owner-controlled token burning to manage supply
- **Enhanced Security:** Multiple security layers including zero address checks and overflow protection
- **Gas Optimization:** Custom error messages reduce transaction costs
- **Front-running Protection:** Built-in validation prevents approve() front-running attacks
- **Ownership Control:** Secure ownership model with renouncement protection

### Technical Advantages

1. **Custom Error Implementation**
   - Uses Solidity 0.8.26 custom errors instead of require strings
   - Significantly reduces gas costs for failed transactions
   - Better developer experience with specific error types

2. **Security Enhancements**
   - Zero address validation on all transfers and approvals
   - Non-zero to non-zero approve protection against front-running
   - Owner renouncement disabled to prevent accidental loss of control
   - Comprehensive input validation on all public functions

3. **Standard Extensions**
   - `increaseAllowance()` and `decreaseAllowance()` functions
   - `getOwner()` function for BEP-20 compliance
   - `burn()` function for supply management
   - Detailed event logging for all operations

4. **Code Quality**
   - Comprehensive documentation with NatSpec comments
   - Clean, readable code structure
   - Follows OpenZeppelin patterns and conventions
   - Single file deployment for transparency

## API Reference

### Core Functions

```javascript
// Get token info
const tokenInfo = await trag.getTokenInfo();

// Check balance
const balance = await trag.getBalance('address');

// Transfer tokens (requires private key)
const result = await trag.transfer('recipient', '100');

// Approve tokens (requires private key)
const result = await trag.approve('spender', '1000');

// Check allowance
const allowance = await trag.getAllowance('owner', 'spender');
```

### Contract Details

- **Decimals**: 6
- **Compiler**: Solidity v0.8.26
- **Security Features**: Custom errors, front-running protection, zero address validation
- **Gas Optimization**: ~51k gas for transfers

For complete API documentation, see [API Reference](docs/API_REFERENCE.md)

## Security

- Verify contract address before transactions
- Source code available on BSCScan for verification
- Bug bounty program - see [SECURITY.md](SECURITY.md)

## Web3 Integration Examples

This repository includes comprehensive Web3 examples for integrating TRAG token functionality into your applications. Examples are provided in multiple programming languages:

### Available Languages

- **[JavaScript/Node.js](examples/javascript/)** - Web3.js integration with full transaction support
- **[Python](examples/python/)** - Web3.py integration with async support
- **[Go](examples/go/)** - go-ethereum integration for high-performance applications
- **[PHP](examples/php/)** - Web3.php integration for web applications

### Common Features

All examples include:

- ✅ Token information retrieval (name, symbol, decimals, total supply)
- ✅ Balance checking for any address
- ✅ Allowance checking between addresses
- ✅ Transfer functionality (requires private key)
- ✅ Approval functionality (requires private key)
- ✅ Error handling and validation
- ✅ BSC network configuration

### Quick Start

Each language directory contains:
- Complete working examples
- Installation instructions
- Usage documentation
- Security best practices

### TRAG Ecosystem Addresses

Official wallet addresses for transparency:

- **TRAG Investor:** `0x7E18Dc9e51965aBC7Db6F9c41Abeb040122400C8`
- **TRAG Team:** `0xB57B316BC2eE445c90e2F4e04DA6D6da1b53cf05`
- **TRAG Rewards:** `0x9a5f69e1bb111789059CaeE4A156FC490e46DC4e`
- **TRAG Treasury:** `0x30B252B746C49651f822d07b12B31B2d73e56af8`
- **TRAG Marketing:** `0xd2839de718cbBe5E13fD4DEFae92010fE6F26383`
- **TRAG Liquidity:** `0x1Cab22De64a3c852AC81fAd9a1590600eCd02325`

## Disclaimer

This is a cryptocurrency token. Please do your own research and understand the risks before investing or trading. The value of cryptocurrency can be volatile and may result in financial loss.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.