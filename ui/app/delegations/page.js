export default function DelegationsPage() {
  const delegations = [
    { id: 'del:001', from: 'did:web:didone.world:humans:alice', to: 'did:web:didone.world:agents:001', scope: 'customer_support', status: 'active' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Delegations</h1>
        <button className="btn">+ Create Delegation</button>
      </div>
      <table className="table">
        <thead><tr><th>From</th><th>To</th><th>Scope</th><th>Status</th></tr></thead>
        <tbody>
          {delegations.map(d => (
            <tr key={d.id}><td><code>{d.from}</code></td><td><code>{d.to}</code></td><td>{d.scope}</td><td style={{color:'#22c55e'}}>{d.status}</td></tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}