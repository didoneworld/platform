"""Wallet SDK"""
import requests
from dataclasses import dataclass
from typing import Dict, List, Optional

@dataclass
class Wallet:
    id: str
    owner_did: str
    created_at: str
    
    @classmethod
    def from_dict(cls, d: Dict) -> 'Wallet':
        return cls(
            id=d.get('id', ''),
            owner_did=d.get('ownerDid', ''),
            created_at=d.get('createdAt', '')
        )

@dataclass
class StoredCredential:
    id: str
    type: List[str]
    issuer: str
    credential_subject: Dict
    
    @classmethod
    def from_dict(cls, d: Dict) -> 'StoredCredential':
        return cls(
            id=d.get('id', ''),
            type=d.get('type', []),
            issuer=d.get('issuer', ''),
            credential_subject=d.get('credentialSubject', {})
        )

class WalletClient:
    """Wallet Client"""
    
    def __init__(self, base_url: str = "http://localhost:3000"):
        self.base_url = base_url.rstrip('/')
    
    @property
    def healthy(self) -> bool:
        resp = requests.get(f"{self.base_url}/health")
        return resp.status_code == 200
    
    def create(self, owner_did: str) -> Wallet:
        """Create wallet"""
        resp = requests.post(f"{self.base_url}/v1/wallets", json={"ownerDid": owner_did})
        resp.raise_for_status()
        return Wallet.from_dict(resp.json())
    
    def get(self, wallet_id: str) -> Wallet:
        """Get wallet"""
        resp = requests.get(f"{self.base_url}/v1/wallets/{wallet_id}")
        resp.raise_for_status()
        return Wallet.from_dict(resp.json())
    
    def add_credential(self, wallet_id: str, credential: Dict) -> Wallet:
        """Add credential to wallet"""
        resp = requests.post(f"{self.base_url}/v1/wallets/{wallet_id}/credentials", json=credential)
        resp.raise_for_status()
        return Wallet.from_dict(resp.json())
    
    def list_credentials(self, wallet_id: str) -> List[StoredCredential]:
        """List credentials"""
        resp = requests.get(f"{self.base_url}/v1/wallets/{wallet_id}/credentials")
        resp.raise_for_status()
        data = resp.json()
        return [StoredCredential.from_dict(c) for c in data.get('credentials', [])]
    
    def present(self, wallet_id: str, credential_id: str, audience: str, challenge: str = None) -> Dict:
        """Create presentation"""
        resp = requests.post(
            f"{self.base_url}/v1/wallets/{wallet_id}/present",
            json={"credentialId": credential_id, "audience": audience, "challenge": challenge or ""}
        )
        resp.raise_for_status()
        return resp.json()
