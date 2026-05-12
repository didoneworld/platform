import Link from 'next/link'

const menu = [
  { label: 'Overview', href: '/', icon: '⌘' },
  { label: 'Identities', href: '/identities', icon: '◉' },
  { label: 'Humans', href: '/humans', icon: '👤' },
  { label: 'Agents', href: '/agents', icon: '🤖' },
  { label: 'APIs', href: '/apis', icon: '⚡' },
  { label: 'Skills', href: '/skills', icon: '⚒' },
  { label: 'Things', href: '/things', icon: '📱' },
  { label: 'Applications', href: '/applications', icon: '📦' },
  { label: 'Organizations', href: '/organizations', icon: '🏢' },
  { label: 'Data Assets', href: '/data-assets', icon: '📄' },
  { label: 'Credentials', href: '/credentials', icon: '🔐' },
  { label: 'Wallets', href: '/wallets', icon: '👛' },
  { label: 'Trust Registry', href: '/trust-registry', icon: '✓' },
  { label: 'Policies', href: '/policies', icon: '📋' },
  { label: 'Delegations', href: '/delegations', icon: '⇤' },
  { label: 'Identity Graph', href: '/graph', icon: '◎' },
  { label: 'Audit Logs', href: '/audit', icon: '📜' },
  { label: 'Compliance', href: '/compliance', icon: '⚖' },
  { label: 'Developers', href: '/developers', icon: '⚡' },
  { label: 'Integrations', href: '/integrations', icon: '🔌' },
  { label: 'Security', href: '/security', icon: '🔒' },
  { label: 'Settings', href: '/settings', icon: '⚙' },
]

export default function Layout({ children }) {
  return (
    <div className="layout">
      <aside className="sidebar">
        <div className="sidebar-header">
          <Link href="/" className="logo">DID One World</Link>
        </div>
        <nav className="nav">
          {menu.map(item => (
            <Link key={item.href} href={item.href} className="nav-item">
              <span className="nav-icon">{item.icon}</span>
              <span>{item.label}</span>
            </Link>
          ))}
        </nav>
      </aside>
      <main className="main">
        <div className="container">
          {children}
        </div>
      </main>
    </div>
  )
}