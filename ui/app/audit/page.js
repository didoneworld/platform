export default function AuditPage() {
  const logs = [
    { id: 1, action: 'agent.created', entity: 'did:web:didone.world:agents:001', actor: 'did:web:didone.world:humans:alice', timestamp: '2026-05-10T10:00:00Z' },
    { id: 2, action: 'credential.issued', entity: 'vc:001', actor: 'did:web:didone.world:org:acme', timestamp: '2026-05-01T09:00:00Z' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Audit Logs</h1>
      </div>
      <table className="table">
        <thead><tr><th>Action</th><th>Entity</th><th>Actor</th><th>Timestamp</th></tr></thead>
        <tbody>
          {logs.map(log => (
            <tr key={log.id}><td>{log.action}</td><td><code>{log.entity}</code></td><td>{log.actor}</td><td>{log.timestamp}</td></tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}