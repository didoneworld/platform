import './styles.css'

const services = [
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
      <header>
        <div className="container header-content">
          <div>
            <h1>DID One World</h1>
            <p>Identity for Everything</p>
          </div>
          <nav>
            <a href="https://github.com/didoneworld">GitHub</a>
            <a href="/docs">Docs</a>
          </nav>
        </div>
      </header>

      <section className="hero">
        <div className="badge">Unified Identity Platform</div>
        <h2>Identity for<br/>Humans, Agents, APIs,<br/>Skills & Things</h2>
        <p className="tagline">
          A unified decentralized identity lifecycle platform combining agent identity, 
          verifiable credentials, and wallet infrastructure.
        </p>
      </section>

      <section className="services">
        <div className="container">
          <h3>Services</h3>
          <div className="grid">
            {services.map(svc => (
              <div key={svc.port} className="card">
                <div className="card-header">
                  <span>{svc.name}</span>
                  <span className="port">:{svc.port}</span>
                </div>
                <p>{svc.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      <section className="stats">
        <div className="stat">
          <div className="stat-value">8</div>
          <div className="stat-label">Identity Types</div>
        </div>
        <div className="stat">
          <div className="stat-value">3</div>
          <div className="stat-label">Repositories</div>
        </div>
        <div className="stat">
          <div className="stat-value">6</div>
          <div className="stat-label">Services</div>
        </div>
      </section>

      <section className="cta">
        <a href="/v1/agents" className="btn-primary">API Console →</a>
        <a href="https://github.com/didoneworld/platform" className="btn-secondary">Documentation</a>
      </section>

      <footer>
        <p>DID One World Platform © 2026 • <a href="https://github.com/didoneworld">GitHub</a></p>
      </footer>
    </>
  )
}