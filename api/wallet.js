const http = require('http');
const url = require('url');
const PORT = process.env.PORT || 3000;

const wallets = {
  'did:web:didone.world:humans:1': { credentials: ['vc:1'], did: 'did:web:didone.world:humans:1' },
  'did:web:didone.world:agents:1': { credentials: ['vc:2'], did: 'did:web:didone.world:agents:1' },
};

const router = (req, res) => {
  const parsed = url.parse(req.url, true);
  res.setHeader('Content-Type', 'application/json');
  res.setHeader('Access-Control-Allow-Origin', '*');
  
  if (parsed.pathname === '/health') {
    res.end(JSON.stringify({ status: 'ok', service: 'wallet' }));
  } else if (parsed.pathname === '/api/wallets') {
    res.end(JSON.stringify({ wallets }));
  } else if (parsed.pathname.startsWith('/api/wallets/')) {
    const did = parsed.pathname.split('/api/wallets/')[1];
    if (wallets[did]) {
      res.end(JSON.stringify(wallets[did]));
    } else {
      res.statusCode = 404;
      res.end(JSON.stringify({ error: 'not found' }));
    }
  } else {
    res.statusCode = 404;
    res.end(JSON.stringify({ error: 'not found' }));
  }
};

require('http').createServer(router).listen(PORT, () => {
  console.log(`Wallet API running on port ${PORT}`);
});
