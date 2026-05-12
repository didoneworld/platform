# DID One World Platform

Unified identity platform combining Agent-DID, verifiable-credential, and idwallet.

## Quick Start

```bash
# Clone with submodules
git clone --recurse-submodules https://github.com/didoneworld/platform.git
cd platform

# Or initialize submodules
git submodule update --init --recursive
```

## Build & Run

```bash
# Build all images
docker-compose build

# Start platform
docker-compose up -d
```

## Architecture

```
┌─────────────┐     ┌──────────────────┐     ┌────────────────┐
│   idwallet  │────▶│    Agent-DID     │────▶│ verifiable-cred│
│  (wallet)   │     │ (identity reg)   │     │ (credentials)  │
└─────────────┘     └──────────────────┘     └────────────────┘
       Port 3000         Port 8000              8001-8004
```

## Services

| Service | Port | Description |
|---------|------|-------------|
| idwallet | 3000 | Frontend wallet UI |
| agent-did | 8000 | Agent identity registry |
| vc-issuer | 8001 | Credential issuer |
| vc-verifier | 8004 | Credential verifier |
| did-resolver | 8002 | DID resolver |
| vc-revocation | 8003 | Status list/revocation |
| gateway | 80/443 | API gateway |

## Environment

```bash
cp .env.example .env
# Edit .env with your secrets
```

## Development

```bash
# Build Agent-DID
cd repos/Agent-DID && docker build -t didoneworld/agent-did .

# Build VC services
cd repos/verifiable-credential
docker build -t didoneworld/vc-issuer -f services/issuer-service/Dockerfile .
docker build -t didoneworld/vc-verifier -f services/verifier-service/Dockerfile .
docker build -t didoneworld/did-resolver -f services/did-resolver/Dockerfile .
docker build -t didoneworld/vc-revocation -f services/revocation-service/Dockerfile .

# Build idwallet
cd repos/idwallet && docker build -t didoneworld/idwallet .
```