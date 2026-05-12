const http = require('http');
const url = require('url');

const PORT = process.env.PORT || 8000;

const identities = [
  { id: 'did:web:didone.world:humans:1', type: 'human', name: 'Admin', status: 'active' },
  { id: 'did:web:didone.world:agents:1', type: 'agent', name: 'Support Agent', status: 'active' },
  { id: 'did:web:didone.world:apis:1', type: 'api', name: 'Payment API', status: 'active' },
];

const router = (req, res) => {
  const parsed = url.parse(req.url, true);
  const path = parsed.pathname;
  
  res.setHeader('Content-Type', 'application/json');
  res.setHeader('Access-Control-Allow-Origin', '*');
  
  if (path === '/health') {
    res.end(JSON.stringify({ status: 'ok', service: 'did-registry' }));
  } else if (path === '/api/identities') {
    res.end(JSON.stringify({ identities }));
  } else if (path === '/api/identities' && req.method === 'POST') {
    let body = '';
    req.on('data', chunk => body += chunk);
    req.on('end', () => {
      const newId = { ...JSON.parse(body), id: `did:web:didone.world:${Date.now()}` };
      identities.push(newId);
      res.end(JSON.stringify({ success: true, id: newId.id }));
    });
  } else if (path.startsWith('/api/identities/')) {
    const id = path.split('/api/identities/')[1];
    const identity = identities.find(i => i.id.includes(id));
    if (identity) {
      res.end(JSON.stringify(identity));
    } else {
      res.statusCode = 404;
      res.end(JSON.stringify({ error: 'not found' }));
    }
  } else {
    res.statusCode = 404;
    res.end(JSON.stringify({ error: 'not found', path }));
  }
};

require('http').createServer(router).listen(PORT, () => {
  console.log(`DID Registry API running on port ${PORT}`);
});
