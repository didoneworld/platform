export default function SettingsPage() {
  return (
    <div>
      <div className="page-header">
        <h1 className="page-title">Settings</h1>
      </div>
      <div style={{ display: 'grid', gap: '1rem', maxWidth: '600px' }}>
        <div style={{ border: '1px solid #222', borderRadius: '0.5rem', padding: '1.5rem' }}>
          <h3 style={{ margin: '0 0 1rem' }}>General</h3>
          <p style={{ color: '#666', fontSize: '0.875rem' }}>Platform name, domain, timezone</p>
        </div>
        <div style={{ border: '1px solid #222', borderRadius: '0.5rem', padding: '1.5rem' }}>
          <h3 style={{ margin: '0 0 1rem' }}>Security</h3>
          <p style={{ color: '#666', fontSize: '0.875rem' }}>API keys, authentication, session settings</p>
        </div>
        <div style={{ border: '1px solid #222', borderRadius: '0.5rem', padding: '1.5rem' }}>
          <h3 style={{ margin: '0 0 1rem' }}>Integrations</h3>
          <p style={{ color: '#666', fontSize: '0.875rem' }}>Webhooks, external services</p>
        </div>
      </div>
    </div>
  )
}