export default function AgentsPage() {
  const agents = [
    { id: 'did:web:didone.world:agents:001', name: 'Support Agent', status: 'active', type: 'agent', owner: 'did:web:didone.world:humans:alice', created: '2026-05-10' },
    { id: 'did:web:didone.world:agents:002', name: 'Sales Agent', status: 'active', type: 'agent', owner: 'did:web:didone.world:humans:bob', created: '2026-05-11' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Agents</h1>
        <button className="btn">+ Create Agent</button>
      </div>
      <table className="table">
        <thead>
          <tr>
            <th>Name</th>
            <th>DID</th>
            <th>Status</th>
            <th>Owner</th>
            <th>Created</th>
          </tr>
        </thead>
        <tbody>
          {agents.map(agent => (
            <tr key={agent.id}>
              <td>{agent.name}</td>
              <td><code>{agent.id}</code></td>
              <td><span style={{ color: '#22c55e' }}>{agent.status}</span></td>
              <td>{agent.owner}</td>
              <td>{agent.created}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}