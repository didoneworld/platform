const http = require('http');
const url = require('url');
const PORT = process.env.PORT || 8001;

const credentials = [
  { id: 'vc:1', type: 'EmployeeCredential', holder: 'did:web:didone.world:humans:1', status: 'valid' },
  { id: 'vc:2', type: 'AgentCredential', holder: 'did:web:didone.world:agents:1', status: 'valid' },
];

const router = (req, res) => {
  const parsed = url.parse(req.url, true);
  res.setHeader('Content-Type', 'application/json');
  res.setHeader('Access-Control-Allow-Origin', '*');
  
  if (parsed.pathname === '/health') {
    res.end(JSON.stringify({ status: 'ok', service: 'vc-issuer' }));
  } else if (parsed.pathname === '/api/credentials') {
    res.end(JSON.stringify({ credentials }));
  } else {
    res.statusCode = 404;
    res.end(JSON.stringify({ error: 'not found' }));
  }
};

require('http').createServer(router).listen(PORT, () => {
  console.log(`VC Issuer running on port ${PORT}`);
});
