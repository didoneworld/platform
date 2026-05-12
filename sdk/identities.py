"""DID & Identity SDK"""
import requests
from dataclasses import dataclass
from typing import Dict, List, Optional, Any

@dataclass
class Identity:
    id: str
    did: str
    type: str
    display_name: str
    owner: str
    controller: str
    status: str
    created_at: str
    
    @classmethod
    def from_dict(cls, d: Dict) -> 'Identity':
        return cls(
            id=d.get('id', ''),
            did=d.get('did', ''),
            type=d.get('type', ''),
            display_name=d.get('displayName', ''),
            owner=d.get('owner', ''),
            controller=d.get('controller', ''),
            status=d.get('status', 'draft'),
            created_at=d.get('createdAt', '')
        )

@dataclass 
class DIDDocument:
    id: str
    controller: str
    verification_method: List[Dict]
    authentication: List[str]
    
    @classmethod
    def from_dict(cls, d: Dict) -> 'DIDDocument':
        return cls(
            id=d.get('id', ''),
            controller=d.get('controller', ''),
            verification_method=d.get('verificationMethod', []),
            authentication=d.get('authentication', [])
        )

class DIDClient:
    """DID Registry Client"""
    
    def __init__(self, base_url: str = "http://localhost:8000"):
        self.base_url = base_url.rstrip('/')
    
    @property
    def healthy(self) -> bool:
        resp = requests.get(f"{self.base_url}/health")
        return resp.status_code == 200
    
    def create(self, type: str, display_name: str, owner: str, controller: str = None) -> Identity:
        """Create new identity"""
        payload = {
            "type": type,
            "displayName": display_name,
            "owner": owner,
            "controller": controller or owner
        }
        resp = requests.post(f"{self.base_url}/v1/identities", json=payload)
        resp.raise_for_status()
        return Identity.from_dict(resp.json())
    
    def resolve(self, did: str) -> DIDDocument:
        """Resolve DID to document"""
        did = did.replace('did:', '')
        resp = requests.get(f"{self.base_url}/v1/did/{did}")
        if resp.status_code == 404:
            return None
        resp.raise_for_status()
        return DIDDocument.from_dict(resp.json())
    
    def list(self, type: str = None) -> List[Identity]:
        """List identities"""
        params = {"type": type} if type else {}
        resp = requests.get(f"{self.base_url}/v1/identities", params=params)
        resp.raise_for_status()
        data = resp.json()
        return [Identity.from_dict(i) for i in data.get('identities', [])]
    
    def get(self, did: str) -> Optional[Identity]:
        """Get identity by DID"""
        did = did.replace('did:', '')
        resp = requests.get(f"{self.base_url}/v1/identities/{did}")
        if resp.status_code == 404:
            return None
        resp.raise_for_status()
        return Identity.from_dict(resp.json())
    
    def update_status(self, did: str, status: str) -> Identity:
        """Update identity status"""
        did = did.replace('did:', '')
        resp = requests.patch(f"{self.base_url}/v1/identities/{did}", json={"status": status})
        resp.raise_for_status()
        return Identity.from_dict(resp.json())
