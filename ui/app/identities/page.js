export default function IdentitiesPage() {
  const identities = [
    { id: 'did:web:didone.world:humans:alice', type: 'human', status: 'active', name: 'Alice Smith' },
    { id: 'did:web:didone.world:agents:001', type: 'agent', status: 'active', name: 'Support Agent' },
    { id: 'did:web:didone.world:apis:payments', type: 'api', status: 'active', name: 'Payments API' },
    { id: 'did:web:didone.world:skills:send-email', type: 'skill', status: 'active', name: 'Send Email' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Identities</h1>
        <button className="btn">+ Create Identity</button>
      </div>
      <table className="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>DID</th>
            <th>Type</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          {identities.map(id => (
            <tr key={id.id}>
              <td>{id.name}</td>
              <td><code>{id.id}</code></td>
              <td>{id.type}</td>
              <td><span style={{ color: '#22c55e' }}>{id.status}</span></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}