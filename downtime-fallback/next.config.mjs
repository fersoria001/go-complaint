/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
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
        ],
      },
};

export default nextConfig;
