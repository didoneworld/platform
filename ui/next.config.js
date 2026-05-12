/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  output: 'export',
  images: { unoptimized: true },
  trailingSlash: true,
  poweredByHeader: false,
  compress: true,
  generateEtags: true,
  httpAgentOptions: { keepAlive: true },
}

module.exports = nextConfig