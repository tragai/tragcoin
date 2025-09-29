const { Web3 } = require('web3');

// TRAG Token Configuration
const TRAG_CONFIG = {
    contractAddress: '0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C',
    abi: [
        {"inputs":[],"stateMutability":"nonpayable","type":"constructor"},
        {"inputs":[],"name":"ApprovalBelowZero","type":"error"},
        {"inputs":[],"name":"BurnExceedsBalance","type":"error"},
        {"inputs":[],"name":"InsufficientAllowance","type":"error"},
        {"inputs":[],"name":"NonZeroToNonZeroApprove","type":"error"},
        {"inputs":[],"name":"RenounceDisabled","type":"error"},
        {"inputs":[],"name":"TransferExceedsBalance","type":"error"},
        {"inputs":[],"name":"ZeroAddress","type":"error"},
        {"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},
        {"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},
        {"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},
        {"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
        {"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
        {"inputs":[{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"burn","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"pure","type":"function"},
        {"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[],"name":"getOwner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
        {"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"pure","type":"function"},
        {"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},
        {"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"view","type":"function"},
        {"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"pure","type":"function"},
        {"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},
        {"inputs":[{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[{"internalType":"address","name":"sender","type":"address"},{"internalType":"address","name":"recipient","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},
        {"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}
    ],
    bscRpcUrl: 'https://bsc-dataseed1.binance.org/'
};

class TragToken {
    constructor(privateKey = null) {
        this.web3 = new Web3(TRAG_CONFIG.bscRpcUrl);
        this.contract = new this.web3.eth.Contract(TRAG_CONFIG.abi, TRAG_CONFIG.contractAddress);

        if (privateKey) {
            this.account = this.web3.eth.accounts.privateKeyToAccount(privateKey);
            this.web3.eth.accounts.wallet.add(this.account);
        }
    }

    // Get token information
    async getTokenInfo() {
        try {
            const [name, symbol, decimals, totalSupply] = await Promise.all([
                this.contract.methods.name().call(),
                this.contract.methods.symbol().call(),
                this.contract.methods.decimals().call(),
                this.contract.methods.totalSupply().call()
            ]);

            return {
                name,
                symbol,
                decimals: parseInt(decimals),
                totalSupply: this.web3.utils.fromWei(totalSupply, 'mwei'), // 6 decimals
                contractAddress: TRAG_CONFIG.contractAddress
            };
        } catch (error) {
            throw new Error(`Failed to get token info: ${error.message}`);
        }
    }

    // Get balance of an address
    async getBalance(address) {
        try {
            const balance = await this.contract.methods.balanceOf(address).call();
            return {
                raw: balance,
                formatted: this.web3.utils.fromWei(balance, 'mwei') // 6 decimals
            };
        } catch (error) {
            throw new Error(`Failed to get balance: ${error.message}`);
        }
    }

    // Get allowance
    async getAllowance(owner, spender) {
        try {
            const allowance = await this.contract.methods.allowance(owner, spender).call();
            return {
                raw: allowance,
                formatted: this.web3.utils.fromWei(allowance, 'mwei')
            };
        } catch (error) {
            throw new Error(`Failed to get allowance: ${error.message}`);
        }
    }

    // Transfer tokens (requires private key)
    async transfer(to, amount) {
        if (!this.account) {
            throw new Error('Private key required for transfers');
        }

        try {
            const amountWei = this.web3.utils.toWei(amount.toString(), 'mwei');
            const tx = this.contract.methods.transfer(to, amountWei);

            const gas = await tx.estimateGas({ from: this.account.address });
            const gasPrice = await this.web3.eth.getGasPrice();

            const result = await tx.send({
                from: this.account.address,
                gas: Math.floor(gas * 1.1),
                gasPrice
            });

            return result;
        } catch (error) {
            throw new Error(`Failed to transfer: ${error.message}`);
        }
    }

    // Approve tokens (requires private key)
    async approve(spender, amount) {
        if (!this.account) {
            throw new Error('Private key required for approval');
        }

        try {
            const amountWei = this.web3.utils.toWei(amount.toString(), 'mwei');
            const tx = this.contract.methods.approve(spender, amountWei);

            const gas = await tx.estimateGas({ from: this.account.address });
            const gasPrice = await this.web3.eth.getGasPrice();

            const result = await tx.send({
                from: this.account.address,
                gas: Math.floor(gas * 1.1),
                gasPrice
            });

            return result;
        } catch (error) {
            throw new Error(`Failed to approve: ${error.message}`);
        }
    }
}

// Usage examples
async function examples() {
    try {
        // Initialize without private key for read-only operations
        const trag = new TragToken();

        // Get token information
        console.log('Token Info:');
        const tokenInfo = await trag.getTokenInfo();
        console.log(tokenInfo);

        // Check balance
        console.log('\nBalance Check:');
        const balance = await trag.getBalance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8');
        console.log(`Balance: ${balance.formatted} TRAG`);

        // For transactions, initialize with private key
        // const tragWithKey = new TragToken('your_private_key_here');
        // const transferResult = await tragWithKey.transfer('recipient_address', '100');

    } catch (error) {
        console.error('Error:', error.message);
    }
}

// Export for use in other modules
module.exports = { TragToken, TRAG_CONFIG };

// Run examples if this file is executed directly
if (require.main === module) {
    examples();
}