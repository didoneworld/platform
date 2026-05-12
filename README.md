# DID One World Platform

> Unified identity platform for humans, agents, APIs, skills, and things

## Services

| Service | Port | UI |
|---------|------|-----|
| [Agent-DID](https://github.com/didoneworld/Agent-DID) | 8000 | Admin Console |
| [verifiable-credential](https://github.com/didoneworld/verifiable-credential) | 8001-4 | Landing |
| [idwallet](https://github.com/didoneworld/idwallet) | 3000 | Wallet |

## Quick Start

```bash
# Clone
git clone https://github.com/didoneworld/platform.git
cd platform

# Configure
cp .env.example .env

# Run
docker-compose up -d
```

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                  Platform UI (port 80/3000)          │
├─────────────────────────────────────────────────────────────┤
│  ┌──────────┐  ┌──────────┐  ┌──────────────────────┐  │
│  │Agent-DID │  │VC Issuer│  │     ID Wallet       │  │
│  │ Admin   │  │        │  │                    │  │
│  └──────────┘  └──────────┘  └──────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## Environment Variables

```bash
# Required
DB_PASSWORD=          # PostgreSQL password
SESSION_SIGNING_SECRET=  # Session signing key

# URLs (auto-configured in compose)
AGENT_DID_URL=http://agent-did:8000
VC_ISSUER_URL=http://vc-issuer:8001
VC_VERIFIER_URL=http://vc-verifier:8004
```

## API Endpoints

### Agent-DID (port 8000)
- `POST /v1/agents` - Create agent
- `GET /v1/agents` - List agents
- `GET /health` - Health check
- `/.well-known/openid-configuration` - OIDC discovery

### Verifiable Credential (ports 8001-8004)
- `POST /v1/credentials/issue` - Issue VC
- `POST /v1/credentials/verify` - Verify VC
- `POST /v1/credentials/revoke` - Revoke VC

## Development

```bash
# Individual services
cd repos/Agent-DID && docker build -t didoneworld/agent-did .
docker build -t didoneworld/vc-issuer -f services/issuer-service/Dockerfile .
```

## Repositories

- [didoneworld/platform](https://github.com/didoneworld/platform) - Platform orchestration
- [didoneworld/Agent-DID](https://github.com/didoneworld/Agent-DID) - Agent identity registry
- [didoneworld/verifiable-credential](https://github.com/didoneworld/verifiable-credential) - VC services
- [didoneworld/idwallet](https://github.com/didoneworld/idwallet) - Wallet SDK