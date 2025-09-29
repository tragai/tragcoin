from web3 import Web3
from web3.middleware import geth_poa_middleware
import json

# TRAG Token Configuration
TRAG_CONFIG = {
    'contract_address': '0x7Cc723dE7fBDb6B06d6628E259e6B8c62673BF1C',
    'abi': [
        {"inputs":[],"stateMutability":"nonpayable","type":"constructor"},
        {"inputs":[],"name":"ApprovalBelowZero","type":"error"},
        {"inputs":[],"name":"BurnExceedsBalance","type":"error"},
        {"inputs":[],"name":"InsufficientAllowance","type":"error"},
        {"inputs":[],"name":"NonZeroToNonZeroApprove","type":"error"},
        {"inputs":[],"name":"RenounceDisabled","type":"error"},
        {"inputs":[],"name":"TransferExceedsBalance","type":"error"},
        {"inputs":[],"name":"ZeroAddress","type":"error"},
        {"anonymous":False,"inputs":[{"indexed":True,"internalType":"address","name":"owner","type":"address"},{"indexed":True,"internalType":"address","name":"spender","type":"address"},{"indexed":False,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},
        {"anonymous":False,"inputs":[{"indexed":True,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":True,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},
        {"anonymous":False,"inputs":[{"indexed":True,"internalType":"address","name":"from","type":"address"},{"indexed":True,"internalType":"address","name":"to","type":"address"},{"indexed":False,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},
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
    'bsc_rpc_url': 'https://bsc-dataseed1.binance.org/'
}

class TragToken:
    def __init__(self, private_key=None):
        # Connect to BSC network
        self.w3 = Web3(Web3.HTTPProvider(TRAG_CONFIG['bsc_rpc_url']))

        # Add PoA middleware for BSC compatibility
        self.w3.middleware_onion.inject(geth_poa_middleware, layer=0)

        # Check connection
        if not self.w3.is_connected():
            raise Exception("Failed to connect to BSC network")

        # Initialize contract
        self.contract = self.w3.eth.contract(
            address=Web3.to_checksum_address(TRAG_CONFIG['contract_address']),
            abi=TRAG_CONFIG['abi']
        )

        # Setup account if private key provided
        self.account = None
        if private_key:
            self.account = self.w3.eth.account.from_key(private_key)

    def get_token_info(self):
        """Get basic token information"""
        try:
            name = self.contract.functions.name().call()
            symbol = self.contract.functions.symbol().call()
            decimals = self.contract.functions.decimals().call()
            total_supply = self.contract.functions.totalSupply().call()

            return {
                'name': name,
                'symbol': symbol,
                'decimals': decimals,
                'total_supply': total_supply / (10 ** decimals),
                'contract_address': TRAG_CONFIG['contract_address']
            }
        except Exception as e:
            raise Exception(f"Failed to get token info: {str(e)}")

    def get_balance(self, address):
        """Get token balance for an address"""
        try:
            balance = self.contract.functions.balanceOf(
                Web3.to_checksum_address(address)
            ).call()

            return {
                'raw': balance,
                'formatted': balance / (10 ** 6)  # 6 decimals
            }
        except Exception as e:
            raise Exception(f"Failed to get balance: {str(e)}")

    def get_allowance(self, owner, spender):
        """Get allowance for spender from owner"""
        try:
            allowance = self.contract.functions.allowance(
                Web3.to_checksum_address(owner),
                Web3.to_checksum_address(spender)
            ).call()

            return {
                'raw': allowance,
                'formatted': allowance / (10 ** 6)
            }
        except Exception as e:
            raise Exception(f"Failed to get allowance: {str(e)}")

    def transfer(self, to_address, amount):
        """Transfer tokens (requires private key)"""
        if not self.account:
            raise Exception("Private key required for transfers")

        try:
            # Convert amount to wei (6 decimals)
            amount_wei = int(amount * (10 ** 6))

            # Build transaction
            transaction = self.contract.functions.transfer(
                Web3.to_checksum_address(to_address),
                amount_wei
            ).build_transaction({
                'from': self.account.address,
                'gas': 100000,
                'gasPrice': self.w3.eth.gas_price,
                'nonce': self.w3.eth.get_transaction_count(self.account.address),
            })

            # Sign and send transaction
            signed_txn = self.w3.eth.account.sign_transaction(transaction, self.account.key)
            tx_hash = self.w3.eth.send_raw_transaction(signed_txn.rawTransaction)

            # Wait for transaction receipt
            receipt = self.w3.eth.wait_for_transaction_receipt(tx_hash)

            return {
                'transaction_hash': tx_hash.hex(),
                'status': receipt['status'],
                'gas_used': receipt['gasUsed']
            }
        except Exception as e:
            raise Exception(f"Failed to transfer: {str(e)}")

    def approve(self, spender_address, amount):
        """Approve tokens for spending (requires private key)"""
        if not self.account:
            raise Exception("Private key required for approval")

        try:
            # Convert amount to wei (6 decimals)
            amount_wei = int(amount * (10 ** 6))

            # Build transaction
            transaction = self.contract.functions.approve(
                Web3.to_checksum_address(spender_address),
                amount_wei
            ).build_transaction({
                'from': self.account.address,
                'gas': 100000,
                'gasPrice': self.w3.eth.gas_price,
                'nonce': self.w3.eth.get_transaction_count(self.account.address),
            })

            # Sign and send transaction
            signed_txn = self.w3.eth.account.sign_transaction(transaction, self.account.key)
            tx_hash = self.w3.eth.send_raw_transaction(signed_txn.rawTransaction)

            # Wait for transaction receipt
            receipt = self.w3.eth.wait_for_transaction_receipt(tx_hash)

            return {
                'transaction_hash': tx_hash.hex(),
                'status': receipt['status'],
                'gas_used': receipt['gasUsed']
            }
        except Exception as e:
            raise Exception(f"Failed to approve: {str(e)}")

# Usage examples
def main():
    try:
        # Initialize without private key for read-only operations
        trag = TragToken()

        # Get token information
        print("Token Info:")
        token_info = trag.get_token_info()
        for key, value in token_info.items():
            print(f"  {key}: {value}")

        # Check balance
        print("\nBalance Check:")
        balance = trag.get_balance('0x742d35cc6564c0532d7e2d9db08d8d37b4b4e6c8')
        print(f"  Balance: {balance['formatted']} TRAG")

        # For transactions, initialize with private key
        # trag_with_key = TragToken('your_private_key_here')
        # transfer_result = trag_with_key.transfer('recipient_address', 100)

    except Exception as e:
        print(f"Error: {str(e)}")

if __name__ == "__main__":
    main()