"""
DID One World SDK

Usage:
    from sdk import DIDClient, WalletClient, VCClient
    
    # Create identity
    client = DIDClient("http://localhost:8000")
    identity = client.create(type="human", display_name="Alice", owner="did:web:example.com:users:alice")
    
    # Resolve DID
    doc = client.resolve(identity.did)
"""

from .identities import DIDClient, DID, Identity
from .wallets import WalletClient, Wallet
from .vc import VCClient, VCIssuer, VCVerifier, VerifiableCredential
