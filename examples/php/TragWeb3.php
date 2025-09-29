<?php

require_once 'vendor/autoload.php';

use Web3\Web3;
use Web3\Contract;
use Web3\Utils;
use Web3\Providers\HttpProvider;
use Web3\RequestManagers\HttpRequestManager;

/**
 * TRAG Token Web3 PHP Client
 *
 * This class provides methods to interact with the TRAG token contract
 * on Binance Smart Chain using Web3 PHP library.
 */
class TragToken
{
    private $web3;
    private $contract;
    private $contractAddress;
    private $contractAbi;
    private $privateKey;
    private $account;

    // TRAG Token Configuration
    const CONTRACT_ADDRESS = '0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C';
    const BSC_RPC_URL = 'https://bsc-dataseed1.binance.org/';
    const TOKEN_DECIMALS = 6;

    // Simplified Contract ABI
    const CONTRACT_ABI = '[
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
    ]';

    /**
     * Constructor
     *
     * @param string|null $privateKey Private key for transactions (optional)
     */
    public function __construct($privateKey = null)
    {
        // Initialize Web3 with BSC provider
        $this->web3 = new Web3(new HttpProvider(new HttpRequestManager(self::BSC_RPC_URL)));

        $this->contractAddress = self::CONTRACT_ADDRESS;
        $this->contractAbi = json_decode(self::CONTRACT_ABI, true);

        // Initialize contract
        $this->contract = new Contract($this->web3->provider, $this->contractAbi);
        $this->contract->at($this->contractAddress);

        // Setup account if private key provided
        if ($privateKey) {
            $this->privateKey = $privateKey;
            // Account setup would require additional Web3 account handling
        }
    }

    /**
     * Get token information
     *
     * @return array Token information
     * @throws Exception
     */
    public function getTokenInfo()
    {
        $result = [];

        // Get name
        $this->contract->call('name', [], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get token name: ' . $err->getMessage());
            }
            $result['name'] = $res[0];
        });

        // Get symbol
        $this->contract->call('symbol', [], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get token symbol: ' . $err->getMessage());
            }
            $result['symbol'] = $res[0];
        });

        // Get decimals
        $this->contract->call('decimals', [], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get token decimals: ' . $err->getMessage());
            }
            $result['decimals'] = intval($res[0]->toString());
        });

        // Get total supply
        $this->contract->call('totalSupply', [], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get total supply: ' . $err->getMessage());
            }
            $totalSupply = $res[0];
            $result['totalSupply'] = $this->fromWei($totalSupply, self::TOKEN_DECIMALS);
            $result['totalSupplyRaw'] = $totalSupply->toString();
        });

        $result['contractAddress'] = $this->contractAddress;

        return $result;
    }

    /**
     * Get token balance for an address
     *
     * @param string $address The address to check
     * @return array Balance information
     * @throws Exception
     */
    public function getBalance($address)
    {
        if (!Utils::isAddress($address)) {
            throw new Exception('Invalid address format');
        }

        $result = [];

        $this->contract->call('balanceOf', [$address], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get balance: ' . $err->getMessage());
            }

            $balance = $res[0];
            $result['raw'] = $balance->toString();
            $result['formatted'] = $this->fromWei($balance, self::TOKEN_DECIMALS);
        });

        return $result;
    }

    /**
     * Get allowance for spender from owner
     *
     * @param string $owner Owner address
     * @param string $spender Spender address
     * @return array Allowance information
     * @throws Exception
     */
    public function getAllowance($owner, $spender)
    {
        if (!Utils::isAddress($owner) || !Utils::isAddress($spender)) {
            throw new Exception('Invalid address format');
        }

        $result = [];

        $this->contract->call('allowance', [$owner, $spender], function ($err, $res) use (&$result) {
            if ($err !== null) {
                throw new Exception('Failed to get allowance: ' . $err->getMessage());
            }

            $allowance = $res[0];
            $result['raw'] = $allowance->toString();
            $result['formatted'] = $this->fromWei($allowance, self::TOKEN_DECIMALS);
        });

        return $result;
    }

    /**
     * Transfer tokens (requires private key)
     *
     * @param string $to Recipient address
     * @param string $amount Amount to transfer
     * @return array Transaction result
     * @throws Exception
     */
    public function transfer($to, $amount)
    {
        if (!$this->privateKey) {
            throw new Exception('Private key required for transfers');
        }

        if (!Utils::isAddress($to)) {
            throw new Exception('Invalid recipient address');
        }

        $amountWei = $this->toWei($amount, self::TOKEN_DECIMALS);

        // This would require additional implementation for transaction signing
        // and sending with the private key
        throw new Exception('Transfer function requires additional Web3 account implementation');
    }

    /**
     * Approve tokens for spending (requires private key)
     *
     * @param string $spender Spender address
     * @param string $amount Amount to approve
     * @return array Transaction result
     * @throws Exception
     */
    public function approve($spender, $amount)
    {
        if (!$this->privateKey) {
            throw new Exception('Private key required for approvals');
        }

        if (!Utils::isAddress($spender)) {
            throw new Exception('Invalid spender address');
        }

        $amountWei = $this->toWei($amount, self::TOKEN_DECIMALS);

        // This would require additional implementation for transaction signing
        // and sending with the private key
        throw new Exception('Approve function requires additional Web3 account implementation');
    }

    /**
     * Convert from wei to human readable format
     *
     * @param mixed $value Wei value
     * @param int $decimals Token decimals
     * @return string Formatted value
     */
    private function fromWei($value, $decimals)
    {
        $divisor = bcpow('10', $decimals);
        return bcdiv($value->toString(), $divisor, $decimals);
    }

    /**
     * Convert from human readable format to wei
     *
     * @param string $value Human readable value
     * @param int $decimals Token decimals
     * @return string Wei value
     */
    private function toWei($value, $decimals)
    {
        $multiplier = bcpow('10', $decimals);
        return bcmul($value, $multiplier);
    }
}

// Usage examples
function main()
{
    try {
        // Initialize without private key for read-only operations
        $trag = new TragToken();

        // Get token information
        echo "Token Info:\n";
        $tokenInfo = $trag->getTokenInfo();
        foreach ($tokenInfo as $key => $value) {
            echo "  $key: $value\n";
        }

        // Check balance
        echo "\nBalance Check:\n";
        $balance = $trag->getBalance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8');
        echo "  Balance: " . $balance['formatted'] . " TRAG\n";

        // For transactions, initialize with private key
        // $tragWithKey = new TragToken('your_private_key_here');
        // $transferResult = $tragWithKey->transfer('recipient_address', '100');

    } catch (Exception $e) {
        echo "Error: " . $e->getMessage() . "\n";
    }
}

// Run examples if this file is executed directly
if (basename(__FILE__) == basename($_SERVER['PHP_SELF'])) {
    main();
}

?>