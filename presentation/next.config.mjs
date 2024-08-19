/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: false,
    transpilePackages: ['three'],
    images: {
        dangerouslyAllowSVG: true,
        remotePatterns: [
          {
            protocol: 'https',
            hostname: 'mirrors.creativecommons.org',
            port: '',
            pathname: '/presskit/icons/**',
          },
          {
            protocol: 'http',
            hostname: 'localhost',
            port: '5170',
            pathname: '/**/**',
          },
          {
            protocol: 'https',
            hostname: 'api.go-complaint.com',
            port: '',
            pathname: '/**/**',
          },
        ],
      },
};

export default nextConfig;
