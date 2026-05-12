'use client'
import './styles.css'

const features = [
  { name: 'Agent ID', port: 8000, desc: 'Agent identity registry & lifecycle' },
  { name: 'VC Issuer', port: 8001, desc: 'Verifiable credential issuance' },
  { name: 'VC Verifier', port: 8004, desc: 'Verifiable credential verification' },
  { name: 'DID Resolver', port: 8002, desc: 'Decentralized identifier resolution' },
  { name: 'VC Revocation', port: 8003, desc: 'Credential status & revocation' },
  { name: 'ID Wallet', port: 3000, desc: 'User wallet & credentials' },
]

export default function Home() {
  return (
    <>
      <div className="hero">
        <div className="badge">Unified Identity Platform</div>
        <h1>Identity for<br/>Humans, Agents, APIs,<br/>Skills & Things</h1>
        <p className="tagline">
          A unified decentralized identity lifecycle platform combining agent identity, 
          verifiable credentials, and wallet infrastructure for the enterprise.
        </p>
        <div className="cta">
          <a href="/identities" className="btn-primary">Get Started →</a>
          <a href="https://github.com/didoneworld/platform" className="btn-secondary">View on GitHub</a>
        </div>
        <p style={{ marginTop: '2rem', fontSize: '0.875rem', color: '#666' }}>
          Trusted by enterprises managing 1000+ identities
        </p>
      </div>

      <section className="features">
        <div className="container">
          <h2>Everything you need for identity management</h2>
          <div className="grid">
            <div className="card">
              <h3>🔐 Universal Identity</h3>
              <p>Manage humans, agents, APIs, skills, and things with a unified decentralized identifier model</p>
            </div>
            <div className="card">
              <h3>🛡️ Verifiable Credentials</h3>
              <p>Issue, verify, and revoke W3C compliant credentials with full lifecycle</p>
            </div>
            <div className="card">
              <h3>⚡ Delegation Engine</h3>
              <p>Securely delegate authority from humans to agents with scoped permissions</p>
            </div>
            <div className="card">
              <h3>📊 Identity Graph</h3>
              <p>Map and visualize relationships between all entities in your organization</p>
            </div>
            <div className="card">
              <h3>📋 Policy Engine</h3>
              <p>Enforce rules across identity creation, verification, and credential issuance</p>
            </div>
            <div className="card">
              <h3>🔒 Enterprise Security</h3>
              <p>Audit logs, compliance reports, and instant revocation for compromised identities</p>
            </div>
          </div>
        </div>
      </section>

      <section className="comparison">
        <div className="container">
          <h2>Why DID One World?</h2>
          <table className="table">
            <thead>
              <tr>
                <th>Feature</th>
                <th>DID One World</th>
                <th>Traditional IAM</th>
                <th>Agent Platforms</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>Human identity</td>
                <td>✅</td>
                <td>✅</td>
                <td>❌</td>
              </tr>
              <tr>
                <td>Agent identity</td>
                <td>✅</td>
                <td>Limited</td>
                <td>✅</td>
              </tr>
              <tr>
                <td>API identity</td>
                <td>✅</td>
                <td>Partial</td>
                <td>❌</td>
              </tr>
              <tr>
                <td>Skill identity</td>
                <td>✅</td>
                <td>❌</td>
                <td>Partial</td>
              </tr>
              <tr>
                <td>Thing/IoT identity</td>
                <td>✅</td>
                <td>Separate</td>
                <td>❌</td>
              </tr>
              <tr>
                <td>Unified wallet</td>
                <td>✅</td>
                <td>Separate</td>
                <td>❌</td>
              </tr>
              <tr>
                <td>Delegation</td>
                <td>✅</td>
                <td>Partial</td>
                <td>Basic</td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      <section className="code-example">
        <div className="container">
          <h2>Simple API</h2>
          <pre>{`# Create an agent
curl -X POST http://localhost:8000/v1/agents \\
  -H "Content-Type: application/json" \\
  -d '{
    "name": "Support Agent",
    "owner": "did:web:didone.world:humans:alice"
  }'

# Issue credential to agent
curl -X POST http://localhost:8001/v1/credentials/issue \\
  -H "Content-Type: application/json" \\
  -d '{
    "subject": "did:web:didone.world:agents:001",
    "type": "AgentCredential",
    "claims": {"role": "support"}
  }'

# Delegate authority
curl -X POST http://localhost:8000/v1/delegations \\
  -H "Content-Type: application/json" \\
  -d '{
    "from": "did:web:didone.world:humans:alice",
    "to": "did:web:didone.world:agents:001",
    "scope": ["read.tickets", "reply.tickets"]
  }'`}</pre>
        </div>
      </section>

      <section className="stats">
        <div className="stat">
          <div className="stat-value">8</div>
          <div className="stat-label">Identity Types</div>
        </div>
        <div className="stat">
          <div className="stat-value">3</div>
          <div className="stat-label">Core Services</div>
        </div>
        <div className="stat">
          <div className="stat-value">20+</div>
          <div className="stat-label">Admin Pages</div>
        </div>
        <div className="stat">
          <div className="stat-value">100%</div>
          <div className="stat-label">Auditable</div>
        </div>
      </section>
    </>
  )
}