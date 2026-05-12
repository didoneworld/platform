export default function PoliciesPage() {
  const policies = [
    { id: 'pol:001', name: 'Finance Approval', description: 'Approve payments above $10,000', status: 'active' },
    { id: 'pol:002', name: 'Agent Trust Level', description: 'Production agents need valid credentials', status: 'active' },
  ]
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Policies</h1>
        <button className="btn">+ Create Policy</button>
      </div>
      <table className="table">
        <thead><tr><th>Name</th><th>Description</th><th>Status</th></tr></thead>
        <tbody>
          {policies.map(p => (
            <tr key={p.id}><td>{p.name}</td><td>{p.description}</td><td style={{color:'#22c55e'}}>{p.status}</td></tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}