"""
DID One World Verifiable Credential SDK

Usage:
    from vc import VCClient, VCIssuer, VCVerifier
    
    # Issue credential
    issuer = VCIssuer("http://localhost:8001")
    vc = issuer.issue(
        holder_did="did:web:example.com:users:alice",
        credential_type="EmployeeCredential",
        claims={"department": "Engineering", "role": "Senior Engineer"}
    )
    
    # Verify credential
    verifier = VCVerifier("http://localhost:8004")
    result = verifier.verify(vc)
    print(f"Verified: {result.verified}")
"""

import requests
from dataclasses import dataclass
from typing import Dict, List, Optional, Any
from datetime import datetime, timedelta

@dataclass
class VerifiableCredential:
    """W3C Verifiable Credential"""
    id: str
    type: List[str]
    issuer: str
    issuance_date: str
    expiration_date: str
    credential_subject: Dict[str, Any]
    proof: Optional[Dict] = None
    
    @classmethod
    def from_dict(cls, data: Dict) -> 'VerifiableCredential':
        return cls(
            id=data.get('id', ''),
            type=data.get('type', []),
            issuer=data.get('issuer', ''),
            issuance_date=data.get('issuanceDateTime', ''),
            expiration_date=data.get('expirationDateTime', ''),
            credential_subject=data.get('credentialSubject', {}),
            proof=data.get('proof')
        )
    
    def to_dict(self) -> Dict:
        return {
            'id': self.id,
            'type': self.type,
            'issuer': self.issuer,
            'issuanceDateTime': self.issuance_date,
            'expirationDateTime': self.expiration_date,
            'credentialSubject': self.credential_subject,
            'proof': self.proof
        }

@dataclass
class VerifyResult:
    """Verification result"""
    verified: bool
    reason: str
    
    @classmethod
    def from_dict(cls, data: Dict) -> 'VerifyResult':
        return cls(
            verified=data.get('verified', False),
            reason=data.get('reason', '')
        )

class VCIssuer:
    """Verifiable Credential Issuer Client"""
    
    def __init__(self, base_url: str = "http://localhost:8001"):
        self.base_url = base_url.rstrip('/')
    
    @property
    def health(self) -> bool:
        """Check issuer health"""
        resp = requests.get(f"{self.base_url}/health")
        return resp.status_code == 200
    
    def issue(
        self,
        holder_did: str,
        credential_type: str,
        claims: Dict[str, Any],
        issuer_did: str = "did:web:didone.world:org:acme",
        expiry: Optional[str] = None
    ) -> VerifiableCredential:
        """
        Issue a verifiable credential
        
        Args:
            holder_did: DID of the credential holder
            credential_type: Type of credential (e.g., "EmployeeCredential")
            claims: Claims to include in the credential
            issuer_did: DID of the issuer
            expiry: Expiration date (RFC3339), defaults to 1 year
        
        Returns:
            VerifiableCredential
        """
        if not expiry:
            expiry = (datetime.utcnow() + timedelta(days=365)).isoformat() + "Z"
        
        payload = {
            "holderDid": holder_did,
            "issuerDid": issuer_did,
            "credentialType": credential_type,
            "claims": claims,
            "expiry": expiry
        }
        
        resp = requests.post(
            f"{self.base_url}/v1/credentials/issue",
            json=payload
        )
        resp.raise_for_status()
        
        data = resp.json()
        return VerifiableCredential.from_dict(data.get('credential', {}))
    
    def revoke(self, credential_id: str) -> bool:
        """Revoke a credential"""
        resp = requests.post(
            f"{self.base_url}/v1/credentials/revoke",
            json={"credentialId": credential_id}
        )
        return resp.status_code == 200
    
    def list(self) -> List[VerifiableCredential]:
        """List all issued credentials"""
        resp = requests.get(f"{self.base_url}/v1/credentials")
        resp.raise_for_status()
        
        data = resp.json()
        return [VerifiableCredential.from_dict(vc) for vc in data.get('credentials', [])]

class VCVerifier:
    """Verifiable Credential Verifier Client"""
    
    def __init__(self, base_url: str = "http://localhost:8004"):
        self.base_url = base_url.rstrip('/')
    
    def verify(self, credential: Any) -> VerifyResult:
        """
        Verify a verifiable credential
        
        Args:
            credential: VerifiableCredential or dict
        
        Returns:
            VerifyResult
        """
        if isinstance(credential, VerifiableCredential):
            credential = credential.to_dict()
        
        payload = {"credential": credential}
        
        resp = requests.post(
            f"{self.base_url}/v1/credentials/verify",
            json=payload
        )
        resp.raise_for_status()
        
        data = resp.json()
        return VerifyResult.from_dict(data)

class VCClient:
    """Unified VC Client (Issuer + Verifier)"""
    
    def __init__(self, issuer_url: str = "http://localhost:8001",
                 verifier_url: str = "http://localhost:8004"):
        self.issuer = VCIssuer(issuer_url)
        self.verifier = VCVerifier(verifier_url)
    
    @property
    def healthy(self) -> bool:
        return self.issuer.health
