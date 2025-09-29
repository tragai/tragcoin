package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TRAG Token Configuration
const (
	ContractAddress = "0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C"
	BSCRPCUrl       = "https://bsc-dataseed1.binance.org/"
	TokenDecimals   = 6
)

// Contract ABI (simplified for examples)
const ContractABI = `[
	{
		"inputs": [],
		"name": "name",
		"outputs": [{"internalType": "string", "name": "", "type": "string"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [{"internalType": "string", "name": "", "type": "string"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [{"internalType": "uint8", "name": "", "type": "uint8"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address", "name": "account", "type": "address"}],
		"name": "balanceOf",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "recipient", "type": "address"},
			{"internalType": "uint256", "name": "amount", "type": "uint256"}
		],
		"name": "transfer",
		"outputs": [{"internalType": "bool", "name": "", "type": "bool"}],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "owner", "type": "address"},
			{"internalType": "address", "name": "spender", "type": "address"}
		],
		"name": "allowance",
		"outputs": [{"internalType": "uint256", "name": "", "type": "uint256"}],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "spender", "type": "address"},
			{"internalType": "uint256", "name": "amount", "type": "uint256"}
		],
		"name": "approve",
		"outputs": [{"internalType": "bool", "name": "", "type": "bool"}],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

// TragToken represents the TRAG token client
type TragToken struct {
	client         *ethclient.Client
	contractAddr   common.Address
	contractABI    abi.ABI
	privateKey     *ecdsa.PrivateKey
	publicAddress  common.Address
}

// TokenInfo holds basic token information
type TokenInfo struct {
	Name            string
	Symbol          string
	Decimals        uint8
	TotalSupply     *big.Int
	ContractAddress string
}

// Balance represents token balance
type Balance struct {
	Raw       *big.Int
	Formatted *big.Float
}

// NewTragToken creates a new TRAG token client
func NewTragToken(privateKeyHex string) (*TragToken, error) {
	// Connect to BSC network
	client, err := ethclient.Dial(BSCRPCUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to BSC network: %v", err)
	}

	// Parse contract address
	contractAddr := common.HexToAddress(ContractAddress)

	// Parse ABI
	contractABI, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	t := &TragToken{
		client:       client,
		contractAddr: contractAddr,
		contractABI:  contractABI,
	}

	// Setup private key if provided
	if privateKeyHex != "" {
		privateKey, err := crypto.HexToECDSA(privateKeyHex)
		if err != nil {
			return nil, fmt.Errorf("invalid private key: %v", err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("failed to get public key")
		}

		t.privateKey = privateKey
		t.publicAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	}

	return t, nil
}

// GetTokenInfo retrieves basic token information
func (t *TragToken) GetTokenInfo() (*TokenInfo, error) {
	var (
		name        string
		symbol      string
		decimals    uint8
		totalSupply *big.Int
	)

	// Call contract methods
	callOpts := &bind.CallOpts{Context: context.Background()}

	// Get name
	result := []interface{}{&name}
	err := t.client.CallContract(context.Background(), t.buildCallMsg("name"), nil)
	if err == nil {
		err = t.contractABI.UnpackIntoInterface(&result, "name", result)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get token name: %v", err)
	}

	// Get symbol
	result = []interface{}{&symbol}
	err = t.client.CallContract(context.Background(), t.buildCallMsg("symbol"), nil)
	if err == nil {
		err = t.contractABI.UnpackIntoInterface(&result, "symbol", result)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get token symbol: %v", err)
	}

	// Get decimals
	result = []interface{}{&decimals}
	err = t.client.CallContract(context.Background(), t.buildCallMsg("decimals"), nil)
	if err == nil {
		err = t.contractABI.UnpackIntoInterface(&result, "decimals", result)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get token decimals: %v", err)
	}

	// Get total supply
	result = []interface{}{&totalSupply}
	err = t.client.CallContract(context.Background(), t.buildCallMsg("totalSupply"), nil)
	if err == nil {
		err = t.contractABI.UnpackIntoInterface(&result, "totalSupply", result)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get total supply: %v", err)
	}

	return &TokenInfo{
		Name:            name,
		Symbol:          symbol,
		Decimals:        decimals,
		TotalSupply:     totalSupply,
		ContractAddress: ContractAddress,
	}, nil
}

// GetBalance retrieves token balance for an address
func (t *TragToken) GetBalance(address string) (*Balance, error) {
	addr := common.HexToAddress(address)

	// Pack method call
	data, err := t.contractABI.Pack("balanceOf", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to pack balanceOf call: %v", err)
	}

	// Make call
	result, err := t.client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &t.contractAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call balanceOf: %v", err)
	}

	// Unpack result
	balance := new(big.Int)
	err = t.contractABI.UnpackIntoInterface(&[]interface{}{balance}, "balanceOf", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack balance: %v", err)
	}

	// Convert to decimal
	decimals := big.NewInt(int64(TokenDecimals))
	divisor := new(big.Int).Exp(big.NewInt(10), decimals, nil)
	formatted := new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt(divisor))

	return &Balance{
		Raw:       balance,
		Formatted: formatted,
	}, nil
}

// GetAllowance retrieves allowance for spender from owner
func (t *TragToken) GetAllowance(owner, spender string) (*Balance, error) {
	ownerAddr := common.HexToAddress(owner)
	spenderAddr := common.HexToAddress(spender)

	// Pack method call
	data, err := t.contractABI.Pack("allowance", ownerAddr, spenderAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to pack allowance call: %v", err)
	}

	// Make call
	result, err := t.client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &t.contractAddr,
		Data: data,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call allowance: %v", err)
	}

	// Unpack result
	allowance := new(big.Int)
	err = t.contractABI.UnpackIntoInterface(&[]interface{}{allowance}, "allowance", result)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack allowance: %v", err)
	}

	// Convert to decimal
	decimals := big.NewInt(int64(TokenDecimals))
	divisor := new(big.Int).Exp(big.NewInt(10), decimals, nil)
	formatted := new(big.Float).Quo(new(big.Float).SetInt(allowance), new(big.Float).SetInt(divisor))

	return &Balance{
		Raw:       allowance,
		Formatted: formatted,
	}, nil
}

// Transfer sends tokens to another address
func (t *TragToken) Transfer(to string, amount *big.Float) (*types.Transaction, error) {
	if t.privateKey == nil {
		return nil, fmt.Errorf("private key required for transfers")
	}

	toAddr := common.HexToAddress(to)

	// Convert amount to wei
	decimals := big.NewInt(int64(TokenDecimals))
	multiplier := new(big.Int).Exp(big.NewInt(10), decimals, nil)
	amountWei := new(big.Int)
	amount.Mul(amount, new(big.Float).SetInt(multiplier)).Int(amountWei)

	// Get nonce
	nonce, err := t.client.PendingNonceAt(context.Background(), t.publicAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := t.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get gas price: %v", err)
	}

	// Create transaction
	auth, err := bind.NewKeyedTransactorWithChainID(t.privateKey, big.NewInt(56)) // BSC Chain ID
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(100000)
	auth.GasPrice = gasPrice

	// Pack transfer data
	data, err := t.contractABI.Pack("transfer", toAddr, amountWei)
	if err != nil {
		return nil, fmt.Errorf("failed to pack transfer data: %v", err)
	}

	// Create and send transaction
	tx := types.NewTransaction(nonce, t.contractAddr, big.NewInt(0), uint64(100000), gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(56)), t.privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}

	err = t.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx, nil
}

// Helper method to build call message
func (t *TragToken) buildCallMsg(method string, args ...interface{}) ethereum.CallMsg {
	data, _ := t.contractABI.Pack(method, args...)
	return ethereum.CallMsg{
		To:   &t.contractAddr,
		Data: data,
	}
}

// Example usage
func main() {
	// Initialize without private key for read-only operations
	trag, err := NewTragToken("")
	if err != nil {
		log.Fatalf("Failed to initialize TRAG client: %v", err)
	}

	// Get token information
	fmt.Println("Token Info:")
	tokenInfo, err := trag.GetTokenInfo()
	if err != nil {
		log.Fatalf("Failed to get token info: %v", err)
	}

	fmt.Printf("  Name: %s\n", tokenInfo.Name)
	fmt.Printf("  Symbol: %s\n", tokenInfo.Symbol)
	fmt.Printf("  Decimals: %d\n", tokenInfo.Decimals)
	fmt.Printf("  Total Supply: %s\n", tokenInfo.TotalSupply.String())
	fmt.Printf("  Contract: %s\n", tokenInfo.ContractAddress)

	// Check balance
	fmt.Println("\nBalance Check:")
	balance, err := trag.GetBalance("0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8")
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	fmt.Printf("  Balance: %s TRAG\n", balance.Formatted.String())

	// For transactions, initialize with private key:
	// tragWithKey, err := NewTragToken("your_private_key_here")
	// tx, err := tragWithKey.Transfer("recipient_address", big.NewFloat(100))
}