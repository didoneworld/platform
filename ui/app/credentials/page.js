export default function CredentialsPage() {
  const credentials = [
    { id: 'vc:001', type: 'EmployeeCredential', holder: 'did:web:didone.world:humans:alice', issuer: 'did:web:didone.world:org:acme', status: 'valid', issued: '2026-05-01' },
    { id: 'vc:002', type: 'AgentCredential', holder: 'did:web:didone.world:agents:001', issuer: 'did:web:didone.world:org:acme', status: 'valid', issued: '2026-05-10' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Credentials</h1>
        <button className="btn">+ Issue Credential</button>
      </div>
      <table className="table">
        <thead>
          <tr>
            <th>Type</th>
            <th>Holder</th>
            <th>Issuer</th>
            <th>Status</th>
            <th>Issued</th>
          </tr>
        </thead>
        <tbody>
          {credentials.map(vc => (
            <tr key={vc.id}>
              <td>{vc.type}</td>
              <td><code>{vc.holder}</code></td>
              <td>{vc.issuer}</td>
              <td><span style={{ color: '#22c55e' }}>{vc.status}</span></td>
              <td>{vc.issued}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}